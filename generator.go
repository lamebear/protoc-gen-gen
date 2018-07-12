package goserve

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

	for name := range g.genFiles {
		d := g.fDescriptions[name]

		for i, s := range d.Service {
			payload := templatePayload{
				Package: "grpc",
				First:   (i == 0),
			}

			_, pkg, _ := d.goPackageOption()
			payload.ProtoPackage = string(pkg)

			payload.Name = s.GetName()
			payload.Imports = serviceImports(s, g.packageWrappers)
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

func serviceImports(sd *descriptor.ServiceDescriptorProto, w map[string]*fileWrapper) []string {
	imports := []string{}
	requiredPackages := make(map[string]bool)
	for _, m := range sd.Method {
		pkg, _ := splitType(m.GetInputType())
		requiredPackages[pkg] = true
		pkg, _ = splitType(m.GetOutputType())
		requiredPackages[pkg] = true
	}

	for pkg := range requiredPackages {
		importDescription := w[pkg]

		path, pkg, _ := importDescription.goPackageOption()
		if strings.HasSuffix(string(path), string(pkg)) {
			imports = append(imports, path.String())
		} else {
			imports = append(imports, fmt.Sprintf("%s %s", pkg, path.String()))
		}
	}

	return imports
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
	_, pkg, _ := packageWrappers[string(protoPackage)].goPackageOption()
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

		ty.ClientStreamed = ty.ClientStreamed || m.GetClientStreaming()

		tmp[iType] = ty

		oType := resolvedType(m.GetOutputType(), w)
		ty, ok = tmp[oType]
		if !ok {
			ty = typeMetadata{
				methodType: oType,
			}
		}

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

func Generate(req *plugin.CodeGeneratorRequest, renderer Renderer) (*plugin.CodeGeneratorResponse, error) {
	g := generator{
		request:         req,
		fDescriptions:   make(map[string]*fileWrapper),
		genFiles:        make(map[string]bool),
		packageWrappers: make(map[string]*fileWrapper),
		renderer:        renderer,
	}

	for _, f := range req.ProtoFile {
		w := &fileWrapper{FileDescriptorProto: f}
		g.fDescriptions[f.GetName()] = w
		g.packageWrappers[f.GetPackage()] = w
	}

	for _, name := range req.FileToGenerate {
		g.genFiles[name] = true
	}

	return g.Generate()
}
