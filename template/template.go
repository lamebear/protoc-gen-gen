//go:generate go-bindata -ignore=\.go -pkg=template -o=bindata.go ./...
package template

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

func defaultTemplate() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}

func NewTemplateRender(templatePath string) (*template.Template, error) {
	if templatePath != "" {
		return nil, errors.New("Rendering a custom template not yet supported")
	}

	tmp, err := template.New("main").Parse(defaultTemplate())
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %v", err)
	}

	return tmp, nil
}
