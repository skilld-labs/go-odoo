package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	odoo "github.com/ahuret/go-odoo"
)

type GeneratorConfiguration struct {
	Odoo          *odoo.Client
	ModelTemplate *template.Template
	FormatModels  bool
	DestFolder    string
}

type Generator struct {
	odoo         *odoo.Client
	tmpl         *template.Template
	destFolder   string
	formatModels bool
}

func NewGenerator(cfg GeneratorConfiguration) *Generator {
	return &Generator{odoo: cfg.Odoo, tmpl: cfg.ModelTemplate, formatModels: cfg.FormatModels, destFolder: cfg.DestFolder}
}

func (g *Generator) HandleModels(models []string) error {
	mm, err := g.getModels(models)
	if err != nil {
		return err
	}
	if err := g.generateModels(mm); err != nil {
		return err
	}
	return nil
}

func (g *Generator) getModels(models []string) ([]*Model, error) {
	if len(models) == 0 {
		var err error
		models, err = g.GetAllModelsName()
		if err != nil {
			return nil, err
		}
	}
	var mm []*Model
	for _, model := range models {
		mfs, err := g.ModelFieldsFromModel(model)
		if err != nil {
			return nil, err
		}
		if len(mfs) == 0 {
			fmt.Printf("error: cannot find fields for model %s, cannot generate it.\n", model)
			continue
		}
		mm = append(mm, NewModel(model, mfs))
	}
	return mm, nil
}

func (g *Generator) GetAllModelsName() ([]string, error) {
	ims, err := g.odoo.FindIrModels(nil, nil)
	if err != nil || ims == nil {
		return []string{}, nil
	}
	var models []string
	for _, im := range *ims {
		models = append(models, im.Model.Get())
	}
	return models, nil
}

func (g *Generator) ModelFieldsFromModel(model string) ([]*ModelField, error) {
	imfs, err := g.odoo.FindIrModelFieldss(odoo.NewCriteria().Add("model", "=", model), nil)
	if err != nil {
		return nil, err
	}
	return g.IrModelFieldsToModelFields(imfs), nil
}

func (g *Generator) IrModelFieldsToModelFields(imfs *odoo.IrModelFieldss) []*ModelField {
	var mfs []*ModelField
	for _, imf := range *imfs {
		mfs = append(mfs, NewModelField(imf.Name.Get(), imf.Ttype.Get().(string)))
	}
	return mfs
}

func (g *Generator) generateModels(models []*Model) error {
	for _, m := range models {
		filePath := g.destFolder + "/" + strings.Replace(m.Name, ".", "_", -1) + ".go"
		output, err := os.Create(filePath)
		if err != nil {
			return err
		}
		if err := g.tmpl.Execute(output, m); err != nil {
			return err
		}
		if g.formatModels {
			if err := exec.Command("gofmt", "-w", filePath).Run(); err != nil {
				return err
			}
		}
		fmt.Printf("%s has been generated\n", filePath)
	}
	return nil
}
