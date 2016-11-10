package main

import (
	"bytes"
	"text/template"
)

func fillTemplate(tmpl string, data interface{}) string {
	var buf bytes.Buffer
	err := template.Must(template.New("").Parse(tmpl)).Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
