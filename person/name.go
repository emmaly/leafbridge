package person

import (
	"bytes"
	"strings"
	"text/template"
)

// Name is a Person's name
type Name struct {
	Family  string
	Middle  string
	Given   string
	Ordinal string
	Prefix  string
	Suffix  string
	Format  NameFormat
}

// NameFormat is a Person's preferred name order/format
type NameFormat uint8

// NameFormat constants
const (
	WesternOrder NameFormat = iota
	EasternOrder
)

var funcMap = template.FuncMap{
	"toUpper": strings.ToUpper,
}

var orderFormat = map[NameFormat]*template.Template{
	WesternOrder: template.Must(template.New("").Funcs(funcMap).Parse("{{.Given}} {{.Family}}{{if .Ordinal}}, {{.Ordinal}}{{end}}{{if .Suffix}}, {{.Suffix}}{{end}}")),
	EasternOrder: template.Must(template.New("").Funcs(funcMap).Parse("{{.Family | toUpper}} {{.Given}}{{if .Ordinal}}, {{.Ordinal}}{{end}}{{if .Suffix}}, {{.Suffix}}{{end}}")),
}

func (n Name) String() string {
	var buf bytes.Buffer
	orderFormat[n.Format].Execute(&buf, n)
	return buf.String()
}
