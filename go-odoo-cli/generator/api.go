package generator

import (
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GenerateAPI(model string) error {
	funcMap := template.FuncMap{
		"funcName": strcase.ToCamel,
	}
	snakeModel := strings.Replace(model, ".", "_", -1)
	tmpl, err := template.New("api.tmpl").Funcs(funcMap).ParseFiles("tmpl/api.tmpl")
	if err != nil {
		return err
	}
	output, err := os.Create("../api/" + snakeModel + "_gen.go")
	if err != nil {
		return err
	}
	err = tmpl.Execute(output, snakeModel)
	if err != nil {
		return err
	}
	return exec.Command("gofmt", "-w", "../api/"+snakeModel+"_gen.go").Run()
}
