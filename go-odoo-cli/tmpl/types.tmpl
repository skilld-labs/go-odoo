package types
{{ if .Time }}
import(
	"time"
)
{{end}}
type {{ .CamelModelName }} struct {
{{ range $key, $field := .Fields }}{{ $field.Name }} {{ $field.Type }} `xmlrpc:"{{ $field.SnakeName }}"`
	{{end}} }

type {{ .CamelModelName }}Nil struct {
	{{ range $key, $field := .Fields }}{{ $field.Name }} {{ $field.NilType }} `xmlrpc:"{{ $field.SnakeName }}"`
	{{end}} }

var {{ .CamelModelName }}Model string = "{{ .ModelName }}"

type {{ .CamelModelName }}s []{{ .CamelModelName }}

type {{ .CamelModelName }}sNil []{{ .CamelModelName }}Nil

func (s *{{ .CamelModelName }}) NilableType_() interface{} {
	return &{{ .CamelModelName }}Nil{}
}

func (ns *{{ .CamelModelName }}Nil) Type_() interface{} {
	s := &{{ .CamelModelName }}{}
	return load(ns, s)
}

func (s *{{ .CamelModelName }}s) NilableType_() interface{} {
	return &{{ .CamelModelName }}sNil{}
}

func (ns *{{ .CamelModelName }}sNil) Type_() interface{} {
	s := &{{ .CamelModelName }}s{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*{{ .CamelModelName }}))
	}
	return s
}
