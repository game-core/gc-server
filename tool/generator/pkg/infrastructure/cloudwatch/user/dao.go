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
	"github.com/game-core/gc-server/internal/errors"
)

const daoTemplate = `
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	"context"
	"time"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"

	{{.Import}}
	"github.com/game-core/gc-server/config/logger"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/pointers"
)

type {{.CamelName}}CloudWatchDao struct {
	ReadCloudWatchConn  *cloudwatchlogs.Client
	WriteCloudWatchConn *cloudwatchlogs.Client
}

func New{{.Name}}CloudWatchDao(conn *logger.CloudWatchHandler) {{.Package}}.{{.Name}}CloudWatchRepository {
	return &{{.CamelName}}CloudWatchDao{
		ReadCloudWatchConn:  conn.User.ReadCloudWatchConn,
		WriteCloudWatchConn: conn.User.WriteCloudWatchConn,
	}
}

{{.Script}}
`

type Dao struct{}

func NewDao() *Dao {
	return &Dao{}
}

// generate 生成する
func (s *Dao) generate(path string, base string) error {
	importCode = ""

	yamlStruct, err := s.getYamlStruct(path)
	if err != nil {
		return err
	}

	domainPath, err := s.getDomainPath(fmt.Sprintf("%s_model.gen.go", changes.UpperCamelToSnake(yamlStruct.Name)))
	if err != nil {
		return err
	}

	if err := NewCloudWatchRepository().generate(path, domainPath); err != nil {
		return err
	}

	outputDir := filepath.Join(base, yamlStruct.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(yamlStruct, s.getOutputFileName(outputDir, changes.UpperCamelToSnake(yamlStruct.Name))); err != nil {
		return err
	}

	return nil
}

// getDomainPath ドメインのpathを取得する関数
func (s *Dao) getDomainPath(name string) (string, error) {
	base := "../../../../../../pkg/domain/model"
	var target string

	if err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() == name {
			target = filepath.Dir(path)
		}
		return nil
	}); err != nil {
		return "", err
	}

	if target == "" {
		return "", errors.NewError("file does not exist")
	}

	importPath := fmt.Sprintf("\"github.com/game-core/gc-server/%s\"", strings.Replace(target, "../../../../../../", "", -1))
	importCode = fmt.Sprintf("%s\n%s", importCode, importPath)

	return target, nil
}

// getYamlStruct yaml構造体を取得する
func (s *Dao) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *Dao) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_dao.gen.go", changes.UpperCamelToSnake(name)))
}

// createOutputFile ファイルを作成する
func (s *Dao) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
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
func (s *Dao) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("daoTemplate").Parse(daoTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"daoTemplate",
		TemplateStruct{
			Name:       yamlStruct.Name,
			Package:    yamlStruct.Package,
			PluralName: changes.SingularToPlural(yamlStruct.Name),
			CamelName:  changes.UpperCamelToCamel(yamlStruct.Name),
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
func (s *Dao) createScript(yamlStruct *YamlStruct) string {
	var methods string

	for _, method := range s.createMethods(yamlStruct) {
		methods = fmt.Sprintf(
			`%s

			%s`,
			methods,
			method,
		)
	}

	return methods
}

// createMethods メソッドを作成する
func (s *Dao) createMethods(yamlStruct *YamlStruct) []string {
	var methods []string

	// Create
	methods = append(methods, s.createCreate(yamlStruct))

	// CreateList
	methods = append(methods, s.createCreateList(yamlStruct))

	// createToCloudWatch
	methods = append(methods, s.createCreateToCloudWatch(yamlStruct))

	// createToFile
	methods = append(methods, s.createCreateToFile(yamlStruct))

	return methods
}

// createCreate Createを作成する
func (s *Dao) createCreate(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`func (s *%sCloudWatchDao) Create(ctx context.Context, now time.Time, level logger.LogLevel, m *%s.%s) {
			timestamp := now.Unix() * 1000
			t := %s
			message := string(logger.SetLogMessage(now, level, t).ToJson())
		
			if os.Getenv("APP_ENV") == "prod" {
				if err := s.creteToCloudWatch(ctx, timestamp, message); err != nil {
					errors.NewMethodErrorLog("appendToFile", err)
				}
			} else if os.Getenv("APP_ENV") == "dev" {
				if err := s.creteToFile("./log/gc_server_user.log", %s, message)); err != nil {
					errors.NewMethodErrorLog("appendToFile", err)
				}
			}
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		yamlStruct.Name,
		s.createTableSetter(yamlStruct),
		"fmt.Sprintf(\"%s %s\\n\", now.Format(time.RFC3339)",
	)
}

// createCreateList CreateListを作成する
func (s *Dao) createCreateList(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`func (s *%sCloudWatchDao) CreateList(ctx context.Context, now time.Time, level logger.LogLevel, ms %s.%s) {
			timestamp := now.Unix() * 1000
			ts := New%s()
			for _, m := range ms {
				t := %s
				ts = append(ts, t)
			}
			message := string(logger.SetLogMessage(now, level, ts).ToJson())
		
			if os.Getenv("APP_ENV") == "prod" {
				if err := s.creteToCloudWatch(ctx, timestamp, message); err != nil {
					errors.NewMethodErrorLog("appendToFile", err)
				}
			} else if os.Getenv("APP_ENV") == "dev" {
				if err := s.creteToFile("./log/gc_server_user.log", %s, message)); err != nil {
					errors.NewMethodErrorLog("appendToFile", err)
				}
			}
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
		s.createTableSetter(yamlStruct),
		"fmt.Sprintf(\"%s %s\\n\", now.Format(time.RFC3339)",
	)
}

// createCreateToCloudWatch createToCloudWatchを作成する
func (s *Dao) createCreateToCloudWatch(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`func (s *%sCloudWatchDao) creteToCloudWatch(ctx context.Context, timestamp int64, message string) error {
			if _, err := s.WriteCloudWatchConn.PutLogEvents(
				ctx,
				&cloudwatchlogs.PutLogEventsInput{
					LogEvents: []types.InputLogEvent{
						{
							Timestamp: &timestamp,
							Message:   &message,
						},
					},
					LogGroupName:  pointers.StringToPointer(os.Getenv("USER_LOG_GROUP_NAME")),
					LogStreamName: pointers.StringToPointer(os.Getenv("USER_LOG_STREAM_NAME")),
				},
			); err != nil {
				errors.NewMethodErrorLog("s.WriteCloudWatchConn.PutLogEvents", err)
			}
		
			return nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
	)
}

// createCreateToFile createToFileを作成する
func (s *Dao) createCreateToFile(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`func (s *%sCloudWatchDao) creteToFile(fileName, message string) error {
			f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			defer func(f *os.File) {
				if err := f.Close(); err != nil {
					errors.NewMethodErrorLog("f.Close", err)
				}
			}(f)
			if _, err := f.WriteString(message); err != nil {
				return err
			}
		
			return nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
	)
}

// checkKeys キーを確認する
func (s *Dao) checkKeys(keys map[string]Structure, name string) bool {
	for _, key := range keys {
		if key.Name == name {
			return true
		}
	}

	return false
}

// checkTimestamp タイムスタンプか確認する
func (s *Dao) checkTimestamp(name string) bool {
	if name != "created_at" && name != "updated_at" {
		return true
	}

	return false
}

// createKeyInterface Interfaceを作成する
func (s *Dao) createKeyInterface(keys map[string]Structure) string {
	var columnStrings []string
	for _, field := range s.getStructures(keys) {
		columnStrings = append(columnStrings, fmt.Sprintf("m.%s", changes.SnakeToUpperCamel(field.Name)))
	}

	return fmt.Sprintf("[]interface{}{%s}", strings.Join(columnStrings, ","))
}

// createUpdateColumn Columnを作成する
func (s *Dao) createUpdateColumn(updates map[string]Structure) string {
	var columnStrings []string
	for _, field := range s.getStructures(updates) {
		columnStrings = append(columnStrings, fmt.Sprintf("\"%s\"", field.Name))
	}

	return strings.Join(columnStrings, ",")
}

// createKeyColumn Columnを作成する
func (s *Dao) createKeyColumn(keys map[string]Structure) string {
	var columnStrings []string
	for _, field := range s.getStructures(keys) {
		columnStrings = append(columnStrings, fmt.Sprintf("{Name: \"%s\"}", field.Name))
	}

	return strings.Join(columnStrings, ",")
}

// createQuery Queryを作成する
func (s *Dao) createQuery(keys map[string]Structure) string {
	var queryStrings []string
	for _, field := range s.getStructures(keys) {
		queryStrings = append(queryStrings, fmt.Sprintf("Where(\"%s = ?\", %s)", field.Name, changes.SnakeToCamel(field.Name)))
	}

	return strings.Join(queryStrings, ".")
}

// createInQuery Queryを作成する
func (s *Dao) createInQuery(keys map[string]Structure) string {
	var queryStrings []string
	for _, field := range s.getStructures(keys) {
		queryStrings = append(queryStrings, field.Name)
	}

	return fmt.Sprintf("Where(\"(%s) IN ?\", ks)", strings.Join(queryStrings, ", "))
}

// createSelect createSelectを作成する
func (s *Dao) createSelect(yamlStruct *YamlStruct) string {
	var paramStrings []string
	for _, field := range s.getStructures(yamlStruct.Structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			paramStrings = append(paramStrings, fmt.Sprintf("\"%s\"", field.Name))
		}
	}

	return fmt.Sprintf(`Select(%s)`, strings.Join(paramStrings, ","))
}

// createModelQuery Queryを作成する
func (s *Dao) createModelQuery(keys map[string]Structure) string {
	var queryStrings []string
	for _, field := range s.getStructures(keys) {
		queryStrings = append(queryStrings, fmt.Sprintf("Where(\"%s = ?\", m.%s)", field.Name, changes.SnakeToUpperCamel(field.Name)))
	}

	return strings.Join(queryStrings, ".")
}

// createParam Paramを作成する
func (s *Dao) createParam(keys map[string]Structure) string {
	var paramStrings []string
	for _, field := range s.getStructures(keys) {
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", changes.SnakeToCamel(field.Name), s.getType(field)))
	}

	return strings.Join(paramStrings, ",")
}

// createModelSetter createModelSetterを作成する
func (s *Dao) createModelSetter(yamlStruct *YamlStruct) string {
	var paramStrings []string
	for _, field := range s.getStructures(yamlStruct.Structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			paramStrings = append(paramStrings, fmt.Sprintf("t.%s,", changes.SnakeToUpperCamel(field.Name)))
		}
	}

	return fmt.Sprintf(
		`%s.Set%s(%s)`,
		yamlStruct.Package,
		yamlStruct.Name,
		strings.Join(paramStrings, ""),
	)
}

// createModelSetters createModelSettersを作成する
func (s *Dao) createModelSetters(yamlStruct *YamlStruct) string {
	var paramStrings []string
	for _, field := range s.getStructures(yamlStruct.Structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			paramStrings = append(paramStrings, fmt.Sprintf("t.%s,", changes.SnakeToUpperCamel(field.Name)))
		}
	}

	return fmt.Sprintf(
		`ms := %s.New%s()
		for _, t := range ts {
			ms = append(ms, %s)
		}`,
		yamlStruct.Package,
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
		fmt.Sprintf(
			`%s.Set%s(%s)`,
			yamlStruct.Package,
			yamlStruct.Name,
			strings.Join(paramStrings, ""),
		),
	)
}

// createTableSetter createTableSetterを作成する
func (s *Dao) createTableSetter(yamlStruct *YamlStruct) string {
	var paramStrings []string
	for _, field := range s.getStructures(yamlStruct.Structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			paramStrings = append(paramStrings, fmt.Sprintf("%s: m.%s,", changes.SnakeToUpperCamel(field.Name), changes.SnakeToUpperCamel(field.Name)))
		}
	}

	return fmt.Sprintf(
		`&%s{
			%s
		}`,
		yamlStruct.Name,
		strings.Join(paramStrings, "\n"),
	)
}

// getStructures フィールド構造体を取得する
func (s *Dao) getStructures(structures map[string]Structure) []*Structure {
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
func (s *Dao) getType(field *Structure) string {
	var result string

	switch field.Type {
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

	if field.Nullable {
		result = fmt.Sprintf("*%s", result)
	}

	return result
}

// getTypeTime timeの型を取得する
func (s *Dao) getTypeTime() string {
	importCode = fmt.Sprintf("%s\n%s", importCode, "\"time\"")
	return "time.Time"
}

// getTypeStructure structureの型を取得する
func (s *Dao) getTypeStructure(fieldName, fieldPackage string) string {
	if fieldPackage != "" {
		importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gc-server/pkg/domain/model/%s\"", fieldPackage))
		return fmt.Sprintf("%s.%s", changes.SnakeToCamel(fieldName), changes.SnakeToUpperCamel(fieldName))
	}

	return changes.SnakeToUpperCamel(fieldName)
}

// getTypeStructures structuresの型を取得する
func (s *Dao) getTypeStructures(fieldName, fieldPackage string) string {
	if fieldPackage != "" {
		importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gc-server/pkg/domain/model/%s\"", fieldPackage))
		return fmt.Sprintf("%s.%s", changes.SnakeToCamel(fieldName), changes.SnakeToUpperCamel(fieldName))
	}

	return changes.SnakeToUpperCamel(fieldName)
}

// getTypeEnum enumの型を取得する
func (s *Dao) getTypeEnum(fieldName, fieldPackage string) string {
	if fieldPackage != "" {
		importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gc-server/pkg/domain/model/%s\"", fieldPackage))
		return fmt.Sprintf("%s.%s", changes.SnakeToCamel(changes.CamelToSnake(changes.Extraction(fieldPackage, "/", 1))), changes.SnakeToUpperCamel(fieldName))
	}

	return changes.SnakeToUpperCamel(fieldName)
}
