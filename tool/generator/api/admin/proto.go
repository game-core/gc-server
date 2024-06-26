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

const protoTemplate = `// {{.Comment}}
syntax = "proto3";

package api.admin;

{{.Script}}
`

type Proto struct{}

func NewProto() *Proto {
	return &Proto{}
}

// generate 生成する
func (s *Proto) generate(file string, base string) error {
	importCode = ""

	yamlStruct, err := s.getYamlStruct(file)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(base, strings.Replace(filepath.Dir(file), "/../../docs/yaml", "", -1))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(file, yamlStruct, s.getOutputFileName(outputDir, filepath.Base(file[:len(file)-len(filepath.Ext(file))]))); err != nil {
		return err
	}

	return nil
}

// getYamlStruct yaml構造体を取得する
func (s *Proto) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *Proto) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.proto", changes.UpperCamelToSnake(name)))
}

// createOutputFile ファイルを作成する
func (s *Proto) createOutputFile(file string, yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := s.createTemplate(file, yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// createTemplate テンプレートを作成する
func (s *Proto) createTemplate(file string, yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("protoTemplate").Parse(protoTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"protoTemplate",
		TemplateStruct{
			Name:    yamlStruct.Name,
			Package: yamlStruct.Package,
			Comment: yamlStruct.Comment,
			Script:  s.createScript(file, yamlStruct),
			Import:  importCode,
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *Proto) createScript(file string, yamlStruct *YamlStruct) string {
	switch {
	case strings.Contains(file, "_request"):
		return s.createStructure(yamlStruct)
	case strings.Contains(file, "_response"):
		return s.createStructure(yamlStruct)
	case strings.Contains(file, "_handler"):
		return s.createService(yamlStruct)
	case strings.Contains(file, "_enum"):
		return s.createEnum(yamlStruct)
	default:
		return s.createStructure(yamlStruct)
	}
}

// createService serviceを作成する
func (s *Proto) createService(yamlStruct *YamlStruct) string {
	var imports []string
	var services []string
	for _, field := range s.getStructure(yamlStruct.Structures) {
		checkReq := true
		req := fmt.Sprintf("import \"%s/%s.proto\";", yamlStruct.Package, changes.UpperCamelToSnake(field.Request))
		for _, m := range imports {
			if m == req {
				checkReq = false
				break
			}
		}
		if checkReq {
			imports = append(imports, req)
		}

		checkRes := true
		res := fmt.Sprintf("import \"%s/%s.proto\";", yamlStruct.Package, changes.UpperCamelToSnake(field.Response))
		for _, m := range imports {
			if m == res {
				checkRes = false
				break
			}
		}
		if checkRes {
			imports = append(imports, res)
		}

		services = append(services, fmt.Sprintf("  rpc %s (%s) returns (%s);", changes.SnakeToUpperCamel(field.Name), field.Request, field.Response))
	}

	script := ""
	if len(imports) == 0 {
		script = fmt.Sprintf(`service %s {
%s
}`,
			yamlStruct.Name,
			strings.Join(services, "\n"),
		)
	} else {
		script = fmt.Sprintf(`%s

service %s {
%s
}`,
			strings.Join(imports, "\n"),
			yamlStruct.Name,
			strings.Join(services, "\n"),
		)
	}

	return script
}

// createStructure structureを作成する
func (s *Proto) createStructure(yamlStruct *YamlStruct) string {
	var imports []string
	var fields []string

	for _, field := range s.getStructure(yamlStruct.Structures) {
		var fe string

		switch field.Type {
		case "structure":
			if field.Package != "" {
				imports = append(imports, fmt.Sprintf("import \"%s/%s.proto\";", field.Package, field.Name))
			} else {
				imports = append(imports, fmt.Sprintf("import \"%s/%s.proto\";", yamlStruct.Package, field.Name))
			}
			fe = fmt.Sprintf("%s %s = %v;", changes.SnakeToUpperCamel(field.Name), field.Name, field.Number)
		case "structures":
			if field.Package != "" {
				imports = append(imports, fmt.Sprintf("import \"%s/%s.proto\";", field.Package, changes.PluralToSingular(field.Name)))
			} else {
				imports = append(imports, fmt.Sprintf("import \"%s/%s.proto\";", yamlStruct.Package, changes.PluralToSingular(field.Name)))
			}
			fe = fmt.Sprintf("repeated %s %s = %v;", changes.SnakeToUpperCamel(changes.PluralToSingular(field.Name)), field.Name, field.Number)
		case "enum":
			imports = append(imports, fmt.Sprintf("import \"%s/%s.proto\";", field.Package, changes.PluralToSingular(field.Name)))
			fe = fmt.Sprintf("%s %s = %v;", changes.SnakeToUpperCamel(field.Name), field.Name, field.Number)
		case "time":
			check := true
			for _, im := range imports {
				if im == "import \"google/protobuf/timestamp.proto\";" {
					check = false
					break
				}
			}
			if check {
				imports = append(imports, "import \"google/protobuf/timestamp.proto\";")
			}
			fe = fmt.Sprintf("google.protobuf.Timestamp %s = %v;", field.Name, field.Number)
		default:
			fe = fmt.Sprintf("%s %s = %v;", field.Type, field.Name, field.Number)
		}

		if field.Nullable {
			fe = "  optional " + fe
		} else {
			fe = "  " + fe
		}

		fields = append(fields, fe)
	}

	script := ""
	if len(imports) == 0 {
		script = fmt.Sprintf(`message %s {
%s
}`,
			yamlStruct.Name,
			strings.Join(fields, "\n"),
		)
	} else {
		script = fmt.Sprintf(`%s

message %s {
%s
}`,
			strings.Join(imports, "\n"),
			yamlStruct.Name,
			strings.Join(fields, "\n"),
		)
	}

	return script
}

// createEnum enumを作成する
func (s *Proto) createEnum(yamlStruct *YamlStruct) string {
	var fields []string

	for _, field := range s.getStructure(yamlStruct.Structures) {
		fields = append(fields, fmt.Sprintf(`  %s = %v;`, field.Name, field.Number))
	}

	return fmt.Sprintf(`enum %s {
%s
}`,
		yamlStruct.Name,
		strings.Join(fields, "\n"),
	)
}

// getStructure フィールド構造体を取得する
func (s *Proto) getStructure(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for _, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:     value.Name,
				Method:   value.Method,
				Request:  value.Request,
				Response: value.Response,
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
