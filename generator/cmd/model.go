package cmd

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

var odooToType = map[string]string{
	"datetime":           "*Time",
	"date":               "*Time",
	"monetary":           "*Float",
	"char":               "*String",
	"many2one":           "*Many2One",
	"many2many":          "*Relation",
	"many2one_reference": "*Many2One",
	"one2many":           "*Relation",
	"integer":            "*Int",
	"boolean":            "*Bool",
	"text":               "*String",
	"selection":          "*Selection",
	"float":              "*Float",
	"binary":             "*String",
	"html":               "*String",
	"reference":          "*String",
}

type model struct {
	Name       string
	StructName string
	VarName    string
	VarsName   string
	Fields     []*modelField
}

type modelField struct {
	Name    string
	VarName string
	Type    string
}

func newModel(name string, mfs []*modelField) *model {
	m := &model{Name: name}
	m.StructName = strcase.ToCamel(m.Name)
	pp := strings.Split(m.Name, ".")
	for _, p := range pp {
		m.VarName += string(p[0])
	}
	if m.VarName == "ids" || m.VarName == "id" || m.VarName == "c" || m.VarName == "if" || m.VarName == "map" {
		m.VarName = strings.ToUpper(m.VarName)
	}
	m.VarsName = m.VarName + "s"
	m.Fields = mfs
	return m
}

func newModelField(name string, odooType string) *modelField {
	mf := &modelField{Name: name}
	mf.VarName = strcase.ToCamel(mf.Name)
	if modelType, ok := odooToType[odooType]; ok {
		mf.Type = modelType
	} else {
		fmt.Printf("warn: cannot find a go-odoo type for odoo type %s\n", odooType)
		mf.Type = "interface{}"
	}
	return mf
}
