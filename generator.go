package gserve

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type generator struct {
	request         *plugin.CodeGeneratorRequest
	fDescriptions   map[string]*fileWrapper
	genFiles        map[string]bool
	packageWrappers map[string]*fileWrapper
	renderer        Renderer
	options         Options
}

type templatePayload struct {
	Imports      []string
	Package      string
	ProtoPackage string
	Name         string
	Methods      []methodPayload
	TypeMetadata []typeMetadata
	First        bool
}

type typeMetadata struct {
	methodType

	IsReceived     bool
	IsSent         bool
	ClientStreamed bool
	ServerStreamed bool
}

type methodType struct {
	Pkg  string
	Name string
}

type methodPayload struct {
	Name            string
	Input           methodType
	Output          methodType
	ClientStreaming bool
	ServerStreaming bool
}

func (g *generator) Generate() (*plugin.CodeGeneratorResponse, error) {
	resp := &plugin.CodeGeneratorResponse{}

	generatedPackage := "grpc"
	if g.options.Package != "" {
		generatedPackage = g.options.Package
	}

	for name := range g.genFiles {
		d := g.fDescriptions[name]
		serviceImport := importString(*d)

		for i, s := range d.Service {
			payload := templatePayload{
				Package: generatedPackage,
				First:   (i == 0),
			}

			_, pkg, _ := d.goPackageOption()
			payload.ProtoPackage = string(pkg)

			payload.Name = s.GetName()
			payload.Imports = serviceImports(s, serviceImport, g.packageWrappers)
			payload.Methods = methodPayloads(d.GetPackage(), s.GetName(), s.Method, g.packageWrappers)
			payload.TypeMetadata = typeMetadatas(s.Method, g.packageWrappers)

			buff := bytes.Buffer{}
			err := g.renderer.Execute(&buff, &payload)
			if err != nil {
				return nil, fmt.Errorf("unable to execute template for %s: %v", name, err)
			}

			c := buff.String()
			fName := fmt.Sprintf("%s.gen.go", strings.ToLower(payload.Name))
			resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
				Name:    &fName,
				Content: &c,
			})
		}
	}

	return resp, nil
}

func serviceImports(sd *descriptor.ServiceDescriptorProto, serviceImport string, w map[string]*fileWrapper) []string {
	importMap := map[string]bool{
		serviceImport: true,
	}
	requiredPackages := make(map[string]bool)
	for _, m := range sd.Method {
		pkg, typ := splitType(m.GetInputType())
		requiredPackages[fmt.Sprintf("%s.%s", pkg, typ)] = true
		pkg, typ = splitType(m.GetOutputType())
		requiredPackages[fmt.Sprintf("%s.%s", pkg, typ)] = true
	}

	for pkg := range requiredPackages {
		importDescription := w[pkg]

		importMap[importString(*importDescription)] = true
	}

	imports := make([]string, 0, len(importMap))
	for importStr := range importMap {
		imports = append(imports, importStr)
	}

	return imports
}

func importString(fw fileWrapper) string {
	path, pkg, _ := fw.goPackageOption()
	if strings.HasSuffix(string(path), string(pkg)) {
		return path.String()
	}
	return fmt.Sprintf("%s %s", pkg, path.String())
}

func methodPayloads(pkg string, svcName string, protoMethods []*descriptor.MethodDescriptorProto, w map[string]*fileWrapper) []methodPayload {
	methods := make([]methodPayload, len(protoMethods))
	for i, m := range protoMethods {
		p := methodPayload{
			Name:            *m.Name,
			Input:           resolvedType(m.GetInputType(), w),
			Output:          resolvedType(m.GetOutputType(), w),
			ClientStreaming: m.GetClientStreaming(),
			ServerStreaming: m.GetServerStreaming(),
		}

		methods[i] = p
	}

	return methods
}

func resolvedType(ts string, packageWrappers map[string]*fileWrapper) methodType {
	protoPackage, ty := splitType(ts)
	_, pkg, _ := packageWrappers[fmt.Sprintf("%s.%s", protoPackage, ty)].goPackageOption()
	return methodType{
		Pkg:  string(pkg),
		Name: ty,
	}
}

func typeMetadatas(methods []*descriptor.MethodDescriptorProto, w map[string]*fileWrapper) []typeMetadata {
	tmp := make(map[methodType]typeMetadata)

	for _, m := range methods {
		iType := resolvedType(m.GetInputType(), w)
		ty, ok := tmp[iType]
		if !ok {
			ty = typeMetadata{
				methodType: iType,
			}
		}

		ty.IsReceived = true
		ty.ClientStreamed = ty.ClientStreamed || m.GetClientStreaming()

		tmp[iType] = ty

		oType := resolvedType(m.GetOutputType(), w)
		ty, ok = tmp[oType]
		if !ok {
			ty = typeMetadata{
				methodType: oType,
			}
		}

		ty.IsSent = true
		ty.ServerStreamed = ty.ServerStreamed || m.GetServerStreaming()

		tmp[oType] = ty
	}

	metadata := make([]typeMetadata, 0, len(tmp))
	for _, t := range tmp {
		metadata = append(metadata, t)
	}

	return metadata
}

func resolveStreamingType(pkg string, svcName string, methodName string) string {
	return fmt.Sprintf("%s.%s_%sServer", pkg, svcName, methodName)
}

func splitType(ts string) (string, string) {
	ts = strings.TrimPrefix(ts, ".")
	i := strings.LastIndex(ts, ".")
	return ts[:i], ts[i+1:]
}

type Renderer interface {
	Execute(io.Writer, interface{}) error
}

func Generate(req *plugin.CodeGeneratorRequest, options Options, renderer Renderer) (*plugin.CodeGeneratorResponse, error) {
	g := generator{
		request:         req,
		fDescriptions:   make(map[string]*fileWrapper),
		genFiles:        make(map[string]bool),
		packageWrappers: make(map[string]*fileWrapper),
		renderer:        renderer,
		options:         options,
	}

	for _, f := range req.ProtoFile {
		w := &fileWrapper{FileDescriptorProto: f}
		g.fDescriptions[f.GetName()] = w

		for _, m := range f.MessageType {
			g.packageWrappers[fmt.Sprintf("%s.%s", f.GetPackage(), m.GetName())] = w
		}

		for _, e := range f.EnumType {
			g.packageWrappers[fmt.Sprintf("%s.%s", f.GetPackage(), e.GetName())] = w
		}
	}

	for _, name := range req.FileToGenerate {
		g.genFiles[name] = true
	}

	return g.Generate()
}
