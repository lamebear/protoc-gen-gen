package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Nais777/goserve"
	"github.com/Nais777/goserve/template"
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

	renderer, err := template.NewTemplateRender("")
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	resp, err := goserve.Generate(req, renderer)
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
