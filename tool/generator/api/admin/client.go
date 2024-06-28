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

const clientTemplate = `
{{.Import}}

{{.Script}}
`

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

// generate 生成する
func (s *Client) generate(file string, base string) error {
	yamlStruct, err := s.getYamlStruct(file)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(base, strings.Replace(filepath.Dir(file), "/../../docs/yaml/api/admin", "/domain/model", -1))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	fileName := s.getOutputFileName(outputDir, filepath.Base(file[:len(file)-len(filepath.Ext(file))]))

	switch {
	case strings.Contains(fileName, "_enum.gen"):
		return nil
	case strings.Contains(fileName, "_handler.gen"):
		return nil
	default:
		if err := s.createOutputFile(yamlStruct, fileName); err != nil {
			return err
		}
		return nil
	}
}

// getYamlStruct yaml構造体を取得する
func (s *Client) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *Client) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.gen.ts", name))
}

// createOutputFile ファイルを作成する
func (s *Client) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
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
func (s *Client) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("templateCode").Parse(clientTemplate)
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
func (s *Client) createScript(yamlStruct *YamlStruct) string {
	var fieldScript []string

	for _, field := range s.getStructure(yamlStruct.Structures) {
		fieldScript = append(fieldScript, fmt.Sprintf("%s: %s;", changes.SnakeToCamel(field.Name), s.getType(yamlStruct, field)))
	}

	return s.createStruct(changes.SnakeToUpperCamel(yamlStruct.Name), strings.Join(fieldScript, "\n"))
}

// getStructure フィールド構造体を取得する
func (s *Client) getStructure(structures map[string]Structure) []*Structure {
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

// createStruct Structを作成する
func (s *Client) createStruct(name string, fieldScript string) string {
	return fmt.Sprintf(
		`export type %s = {
			%s
		}`,
		name,
		fieldScript,
	)
}

// getType 型を取得する
func (s *Client) getType(yamlStruct *YamlStruct, field *Structure) string {
	var result string

	switch field.Type {
	case "int":
	case "int32":
	case "int64":
		result = "number"
	case "time":
		result = s.getTypeTime()
	case "structure":
		result = s.getTypeStructure(field.Name, field.Package)
	case "structures":
		result = s.getTypeStructures(field.Name, field.Package)
	case "enum":
		result = s.getTypeEnum(field.Name, field.Package)
	default:
		result = field.Type
	}

	return result
}

// getTypeTime timeの型を取得する
func (s *Client) getTypeTime() string {
	importCode = fmt.Sprintf("%s\n%s", importCode, "\"time\"")
	return "string"
}

// getTypeStructure structureの型を取得する
func (s *Client) getTypeStructure(fieldName, fieldPackage string) string {
	importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("import type {%s} from \"~/pkg/domain/model/%s/%s.gen\"", changes.SnakeToUpperCamel(fieldName), fieldPackage, fieldName))
	return changes.SnakeToUpperCamel(fieldName)
}

// getTypeStructure structureの型を取得する
func (s *Client) getTypeStructures(fieldName, fieldPackage string) string {
	importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("import type {%s} from \"~/pkg/domain/model/%s/%s.gen\"", changes.SnakeToUpperCamel(fieldName), fieldPackage, fieldName))
	return fmt.Sprintf("%s[]", changes.SnakeToUpperCamel(fieldName))
}

// getTypeEnum enumの型を取得する
func (s *Client) getTypeEnum(fieldName, fieldPackage string) string {
	importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("import type {%s} from \"~/pkg/domain/model/%s/%s.gen\"", changes.SnakeToUpperCamel(fieldName), fieldPackage, fieldName))
	return changes.SnakeToUpperCamel(fieldName)
}
