package gserve

import (
	"path"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type fileWrapper struct {
	*descriptor.FileDescriptorProto
}

func (d fileWrapper) goPackageOption() (impPath GoImportPath, pkg GoPackageName, ok bool) {
	opt := d.GetOptions().GetGoPackage()
	if opt == "" {
		return "", "", false
	}
	// A semicolon-delimited suffix delimits the import path and package name.
	sc := strings.Index(opt, ";")
	if sc >= 0 {
		return GoImportPath(opt[:sc]), cleanPackageName(opt[sc+1:]), true
	}
	// The presence of a slash implies there's an import path.
	slash := strings.LastIndex(opt, "/")
	if slash >= 0 {
		return GoImportPath(opt), cleanPackageName(opt[slash+1:]), true
	}
	return "", cleanPackageName(opt), true
}

func (f fileWrapper) destinationFileName() string {
	name := *f.Name
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	name += ".gen.go"

	if impPath, _, ok := f.goPackageOption(); ok && impPath != "" {
		// Replace the existing dirname with the declared import path.
		_, name = path.Split(name)
		name = path.Join(string(impPath), name)
		return name
	}

	return name
}
