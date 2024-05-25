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

const templateCode = `
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	{{.Import}}
)

type {{.PluralName}} []*{{.Name}}

{{.Script}}
`

type Model struct{}

func NewModel() *Model {
	return &Model{}
}

// generate 生成する
func (s *Model) generate(file string, base string) error {
	yamlStruct, err := s.getYamlStruct(file)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(base, strings.Replace(filepath.Dir(file), "/../../docs/yaml", "", -1))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	fileName := s.getOutputFileName(outputDir, filepath.Base(file[:len(file)-len(filepath.Ext(file))]))

	switch {
	case strings.Contains(fileName, "_enum"):
		return nil
	default:
		if err := s.createOutputFile(yamlStruct, fileName); err != nil {
			return err
		}
		return nil
	}
}

// getYamlStruct yaml構造体を取得する
func (s *Model) getYamlStruct(file string) (*YamlStruct, error) {
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

// getOutputFileName ファイル名を取得する
func (s *Model) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_model.gen.go", changes.UpperCamelToSnake(name)))
}

// createOutputFile ファイルを作成する
func (s *Model) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := s.createTemplate(yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// createTemplate テンプレートを作成する
func (s *Model) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("templateCode").Parse(templateCode)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"templateCode",
		TemplateStruct{
			Name:       yamlStruct.Name,
			Package:    yamlStruct.Package,
			PluralName: changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
			Comment:    yamlStruct.Comment,
			Script:     s.createScript(yamlStruct),
			Import:     importCode,
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *Model) createScript(yamlStruct *YamlStruct) string {
	var fieldScript []string
	var paramScript []string
	var returnScript []string

	for _, field := range s.getStructure(yamlStruct.Structures) {
		fieldScript = append(fieldScript, fmt.Sprintf("%s %s", changes.SnakeToUpperCamel(field.Name), s.getType(yamlStruct, field)))
		paramScript = append(paramScript, fmt.Sprintf("%s %s", changes.SnakeToCamel(field.Name), s.getType(yamlStruct, field)))
		returnScript = append(returnScript, fmt.Sprintf("%s: %s,", changes.SnakeToUpperCamel(field.Name), changes.SnakeToCamel(field.Name)))
	}

	return fmt.Sprintf(
		`%s

		%s

		%s`,
		s.createStruct(yamlStruct.Name, strings.Join(fieldScript, "\n")),
		s.createNew(yamlStruct.Name, changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name)))),
		s.createSetter(yamlStruct.Name, strings.Join(paramScript, ","), strings.Join(returnScript, "\n")),
	)
}

// createStruct Structを作成する
func (s *Model) createStruct(name string, fieldScript string) string {
	return fmt.Sprintf(
		`type %s struct {
			%s
		}`,
		name,
		fieldScript,
	)
}

// createNew Newを作成する
func (s *Model) createNew(name, pluralName string) string {
	return fmt.Sprintf(
		`func New%s() *%s {
			return &%s{}
		}

		func New%s() %s {
			return %s{}
		}`,
		name,
		name,
		name,
		pluralName,
		pluralName,
		pluralName,
	)
}

// createSetter Setterを作成する
func (s *Model) createSetter(name, paramScript, returnScript string) string {
	return fmt.Sprintf(
		`func Set%s(%s) *%s {
			return &%s{
				%s
			}
		}`,
		name,
		paramScript,
		name,
		name,
		returnScript,
	)
}

// getStructure フィールド構造体を取得する
func (s *Model) getStructure(structures map[string]Structure) []*Structure {
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
func (s *Model) getType(yamlStruct *YamlStruct, field *Structure) string {
	var result string

	switch field.Type {
	case "time":
		result = s.getTypeTime()
	case "structure":
		result = s.getTypeStructure(field.Name, field.Package, yamlStruct.Package)
	case "structures":
		result = s.getTypeStructures(field.Name, field.Package, yamlStruct.Package)
	case "enum":
		result = s.getTypeEnum(field.Name, field.Package, yamlStruct.Package)
	default:
		result = field.Type
	}

	if field.Nullable {
		result = fmt.Sprintf("*%s", result)
	}

	return result
}

// getTypeTime timeの型を取得する
func (s *Model) getTypeTime() string {
	importCode = fmt.Sprintf("%s\n%s", importCode, "\"time\"")
	return "time.Time"
}

// getTypeStructure structureの型を取得する
func (s *Model) getTypeStructure(fieldName, fieldPackage, structPackage string) string {
	return s.createType(strings.Split(fieldPackage, "/"), fieldName, fieldPackage, structPackage)
}

// getTypeStructures structuresの型を取得する
func (s *Model) getTypeStructures(fieldName, fieldPackage, structPackage string) string {
	return s.createType(strings.Split(fieldPackage, "/"), fieldName, fieldPackage, structPackage)
}

// getTypeEnum enumの型を取得する
func (s *Model) getTypeEnum(fieldName, fieldPackage, structPackage string) string {
	return s.createType(strings.Split(fieldPackage, "/"), fieldName, fieldPackage, structPackage)
}

// createType 型を生成する
func (s *Model) createType(parts []string, fieldName, fieldPackage, structPackage string) string {
	if len(parts) > 1 {
		if parts[1] != structPackage {
			importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gc-server/pkg/domain/model/%s\"", fieldPackage))
			return fmt.Sprintf("%s.%s", parts[1], changes.SnakeToUpperCamel(fieldName))
		}
	}

	if len(parts) <= 1 {
		if parts[0] != structPackage {
			importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gc-server/pkg/domain/model/%s\"", fieldPackage))
			return fmt.Sprintf("%s.%s", parts[0], changes.SnakeToUpperCamel(fieldName))
		}
	}

	return changes.SnakeToUpperCamel(fieldName)
}
