package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/game-core/gc-server/internal/changes"
)

type CloudWatchRepository struct{}

func NewCloudWatchRepository() *CloudWatchRepository {
	return &CloudWatchRepository{}
}

const repositoryTemplate = `
{{.Mock}}
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	"context"
	"time"

	"github.com/game-core/gc-server/pkg/domain/enum"
	"github.com/game-core/gc-server/config/logger"
)

{{.Script}}
`

// generate 生成する
func (s *CloudWatchRepository) generate(path string, base string) error {
	yamlStruct, err := s.getYamlStruct(path)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(base, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(yamlStruct, s.getOutputFileName(base, changes.UpperCamelToSnake(yamlStruct.Name))); err != nil {
		return err
	}

	return nil
}

// getYamlStruct yaml構造体を取得する
func (s *CloudWatchRepository) getYamlStruct(file string) (*YamlStruct, error) {
	yamlData, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var yamlStruct YamlStruct
	if err := yaml.Unmarshal(yamlData, &yamlStruct); err != nil {
		return nil, err
	}

	return &yamlStruct, nil
}

// createOutputFile ファイルを作成する
func (s *CloudWatchRepository) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := s.createTemplate(yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// getOutputFileName ファイル名を取得する
func (s *CloudWatchRepository) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_mysql_repository.gen.go", changes.UpperCamelToSnake(name)))
}

// createTemplate テンプレートを作成する
func (s *CloudWatchRepository) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("repositoryTemplate").Parse(repositoryTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"repositoryTemplate",
		TemplateStruct{
			Name:       yamlStruct.Name,
			Package:    yamlStruct.Package,
			PluralName: changes.SingularToPlural(yamlStruct.Name),
			CamelName:  changes.UpperCamelToCamel(yamlStruct.Name),
			Comment:    yamlStruct.Comment,
			Script:     s.createScript(yamlStruct),
			Import:     importCode,
			Mock:       createMock(yamlStruct),
		},
	); err != nil {
		return err
	}

	return nil
}

// createMock mockを作成する
func createMock(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		"//go:generate mockgen -source=./%s_mysql_repository.gen.go -destination=./%s_mysql_repository_mock.gen.go -package=%s",
		changes.UpperCamelToSnake(yamlStruct.Name),
		changes.UpperCamelToSnake(yamlStruct.Name),
		changes.UpperCamelToCamel(yamlStruct.Package),
	)
}

// createScript スクリプトを作成する
func (s *CloudWatchRepository) createScript(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`type %sCloudWatchRepository interface {
			%s
		}`,
		yamlStruct.Name,
		strings.Join(s.createMethods(yamlStruct), "\n"),
	)
}

// createMethods メソッドを作成する
func (s *CloudWatchRepository) createMethods(yamlStruct *YamlStruct) []string {
	var methods []string

	// Create
	methods = append(methods, s.createCreate(yamlStruct))

	// CreateList
	methods = append(methods, s.createCreateList(yamlStruct))

	return methods
}

// createCreate Createを作成する
func (s *CloudWatchRepository) createCreate(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`Create(ctx context.Context, now time.Time, level logger.LogLevel, m *%s)`,
		yamlStruct.Name,
	)
}

// createCreateList CreateListを作成する
func (s *CloudWatchRepository) createCreateList(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`CreateList(ctx context.Context, now time.Time, level logger.LogLevel, ms %s)`,
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
	)
}

// createParam Paramを作成する
func (s *CloudWatchRepository) createParam(structPackage string, keys map[string]Structure) string {
	var paramStrings []string
	for _, field := range s.getStructures(keys) {
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", changes.SnakeToCamel(field.Name), s.getType(structPackage, field)))
	}

	return strings.Join(paramStrings, ",")
}

// getStructures フィールド構造体を取得する
func (s *CloudWatchRepository) getStructures(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for _, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:     value.Name,
				Type:     value.Type,
				Package:  value.Package,
				Nullable: value.Nullable,
				Number:   value.Number,
				Comment:  value.Comment,
			},
		)
	}

	sort.Slice(sortStructures, func(i, j int) bool {
		return sortStructures[i].Number < sortStructures[j].Number
	})

	return sortStructures
}

// getType 型を取得する
func (s *CloudWatchRepository) getType(structPackage string, field *Structure) string {
	var result string

	switch field.Type {
	case "time":
		result = s.getTypeTime()
	case "structure":
		result = s.getTypeStructure(field.Name, field.Package, structPackage)
	case "structures":
		result = s.getTypeStructures(field.Name, field.Package, structPackage)
	case "enum":
		result = s.getTypeEnum(field.Name, field.Package, structPackage)
	default:
		result = field.Type
	}

	if field.Nullable {
		result = fmt.Sprintf("*%s", result)
	}

	return result
}

// getTypeTime timeの型を取得する
func (s *CloudWatchRepository) getTypeTime() string {
	return "time.Time"
}

// getTypeStructure structureの型を取得する
func (s *CloudWatchRepository) getTypeStructure(fieldName, fieldPackage string, structPackage string) string {
	if changes.Extraction(fieldPackage, "/", 1) != structPackage {
		return fmt.Sprintf("%s.%s", changes.SnakeToCamel(fieldName), changes.SnakeToUpperCamel(fieldName))
	}

	return changes.SnakeToUpperCamel(fieldName)
}

// getTypeStructures structuresの型を取得する
func (s *CloudWatchRepository) getTypeStructures(fieldName, fieldPackage string, structPackage string) string {
	if changes.Extraction(fieldPackage, "/", 1) != structPackage {
		return fmt.Sprintf("%s.%s", changes.SnakeToCamel(fieldName), changes.SnakeToUpperCamel(changes.SingularToPlural(fieldName)))
	}

	return changes.SnakeToUpperCamel(fieldName)
}

// getTypeEnum enumの型を取得する
func (s *CloudWatchRepository) getTypeEnum(fieldName, fieldPackage, structPackage string) string {
	if changes.Extraction(fieldPackage, "/", 1) != structPackage {
		importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gc-server/pkg/domain/model/%s\"", fieldPackage))
		return fmt.Sprintf("%s.%s", changes.SnakeToCamel(changes.CamelToSnake(changes.Extraction(fieldPackage, "/", 1))), changes.SnakeToUpperCamel(fieldName))
	}

	return changes.SnakeToUpperCamel(fieldName)
}
