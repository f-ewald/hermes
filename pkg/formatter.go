package hermes

import (
	"bytes"
	"encoding/json"
	"github.com/f-ewald/hermes/templates"
	"gopkg.in/yaml.v2"
	"html/template"
)

type Formatter interface {
	Format(interface{}) ([]byte, error)
}

type TextFormatter struct {
	tpl string
}

func NewTextFormatter(tpl string) Formatter {
	return &TextFormatter{tpl: tpl}
}

func (formatter *TextFormatter) Format(a interface{}) ([]byte, error) {
	b, err := templates.Templates.ReadFile("statistics.tpl")
	if err != nil {
		return nil, err
	}
	temp, err := template.New("tpl").Parse(string(b))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = temp.Execute(buf, a)
	if err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}

type JsonFormatter struct{}

func (formatter *JsonFormatter) Format(a interface{}) ([]byte, error) {
	return json.Marshal(a)
}

type YamlFormatter struct{}

func (formatter *YamlFormatter) Format(a interface{}) ([]byte, error) {
	return yaml.Marshal(a)
}
