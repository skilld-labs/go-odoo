package generator

import (
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var convertTypes = map[string]string{
	"datetime":  "time.Time",
	"date":      "time.Time",
	"monetary":  "float64",
	"char":      "string",
	"many2one":  "Many2One",
	"many2many": "[]int64",
	"one2many":  "[]int64",
	"integer":   "int64",
	"boolean":   "bool",
	"text":      "string",
	"selection": "interface{}",
	"float":     "float64",
	"binary":    "string",
	"html":      "string",
	"reference": "string",
}

var convertNilTypes = map[string]string{
	"datetime":  "interface{}",
	"date":      "interface{}",
	"monetary":  "interface{}",
	"char":      "interface{}",
	"many2one":  "interface{}",
	"many2many": "interface{}",
	"one2many":  "interface{}",
	"integer":   "interface{}",
	"boolean":   "bool",
	"text":      "interface{}",
	"selection": "interface{}",
	"float":     "interface{}",
	"binary":    "interface{}",
	"html":      "interface{}",
	"reference": "interface{}",
}

type ModelType struct {
	ModelName      string
	CamelModelName string
	Fields         []Field
	Time           bool
}

type Field struct {
	Name      string
	SnakeName string
	Type      string
	NilType   string
}

func GenerateTypes(model string, fields map[string]string) error {
	snakeModel := strings.Replace(model, ".", "_", -1)
	modelType := ModelType{ModelName: model, CamelModelName: strcase.ToCamel(snakeModel)}
	for fieldName, fieldType := range fields {
		convertType := convertTypes[fieldType]
		if convertType == "time.Time" {
			modelType.Time = true
		}
		f := Field{Name: strcase.ToCamel(fieldName), SnakeName: fieldName, Type: convertType, NilType: convertNilTypes[fieldType]}
		modelType.Fields = append(modelType.Fields, f)
	}
	tmpl, err := template.New("types.tmpl").ParseFiles("tmpl/types.tmpl")
	if err != nil {
		return err
	}
	output, err := os.Create("../types/" + snakeModel + "_gen.go")
	if err != nil {
		return err
	}
	err = tmpl.Execute(output, modelType)
	if err != nil {
		return err
	}
	return exec.Command("gofmt", "-w", "../types/"+snakeModel+"_gen.go").Run()
}
