package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Nais777/protoc-gen-gen/template"
	"github.com/gogo/protobuf/proto"
	p "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	req := &p.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, req); err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	options, err := ParseOptions(req.Parameter)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	renderer, err := template.NewTemplateRender(options.TemplatePath)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	resp, err := Generate(req, options, renderer)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	data, err = proto.Marshal(resp)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}
