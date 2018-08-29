package template

import (
	"fmt"
	"text/template"
)

func NewTemplateRender(templatePath string) (*template.Template, error) {
	tmp, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %v", err)
	}

	return tmp, nil
}
