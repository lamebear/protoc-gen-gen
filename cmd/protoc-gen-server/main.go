package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Nais777/gserve"
	"github.com/Nais777/gserve/template"
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

	options, err := gserve.ParseOptions(req.Parameter)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	renderer, err := template.NewTemplateRender(options.TemplatePath)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	resp, err := gserve.Generate(req, options, renderer)
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
