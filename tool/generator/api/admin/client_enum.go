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

const clientEnumTemplate = `{{.Script}}
`

type ClientEnum struct{}

func NewClientEnum() *ClientEnum {
	return &ClientEnum{}
}

// generate 生成する
func (s *ClientEnum) generate(file string, base string) error {
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
	case strings.Contains(fileName, "Enum.gen"):
		if err := s.createOutputFile(yamlStruct, fileName); err != nil {
			return err
		}
		return nil
	default:
		return nil
	}
}

// getYamlStruct yaml構造体を取得する
func (s *ClientEnum) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *ClientEnum) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
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
func (s *ClientEnum) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.gen.ts", changes.SnakeToUpperCamel(name)))
}

// createTemplate テンプレートを作成する
func (s *ClientEnum) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("clientEnumTemplate").Parse(clientEnumTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"clientEnumTemplate",
		TemplateStruct{
			Name:    yamlStruct.Name,
			Package: yamlStruct.Package,
			Comment: yamlStruct.Comment,
			Script:  s.createScript(yamlStruct),
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *ClientEnum) createScript(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`export enum %s {
			%s
		}`,
		changes.SnakeToUpperCamel(yamlStruct.Name),
		strings.Join(s.createConstants(yamlStruct), "\n"),
	)
}

// createConstants
func (s *ClientEnum) createConstants(yamlStruct *YamlStruct) []string {
	var constants []string
	for _, constant := range s.getStructure(yamlStruct.Structures) {
		constants = append(
			constants,
			fmt.Sprintf(
				"%s = %d,",
				changes.SnakeToUpperCamel(constant.Name),
				constant.Number,
			),
		)
	}

	return constants
}

// getStructure フィールド構造体を取得する
func (s *ClientEnum) getStructure(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for _, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:    value.Name,
				Number:  value.Number,
				Comment: value.Comment,
			},
		)
	}

	sort.Slice(sortStructures, func(i, j int) bool {
		return sortStructures[i].Number < sortStructures[j].Number
	})

	return sortStructures
}
