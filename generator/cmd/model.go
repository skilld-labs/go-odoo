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

type Model struct {
	Name       string
	StructName string
	VarName    string
	VarsName   string
	Fields     []*ModelField
}

type ModelField struct {
	Name    string
	VarName string
	Type    string
}

func NewModel(name string, mfs []*ModelField) *Model {
	m := &Model{Name: name}
	m.StructName = strcase.ToCamel(m.Name)
	pp := strings.Split(m.Name, ".")
	for _, p := range pp {
		m.VarName += string(p[0])
	}
	if m.VarName == "ids" || m.VarName == "id" || m.VarName == "c" || m.VarName == "if" {
		m.VarName = strings.ToUpper(m.VarName)
	}
	m.VarsName = m.VarName + "s"
	m.Fields = mfs
	return m
}

func NewModelField(name string, odooType string) *ModelField {
	mf := &ModelField{Name: name}
	mf.VarName = strcase.ToCamel(mf.Name)
	if modelType, ok := odooToType[odooType]; ok {
		mf.Type = modelType
	} else {
		fmt.Printf("warn: cannot find a go-odoo type for odoo type %s\n", odooType)
		mf.Type = "interface{}"
	}
	return mf
}
