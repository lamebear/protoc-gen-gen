# gServe

gServe aims to ease the creation of go gRPC servers (and eventualy clients) by generating boilerplate code from protobuf files.

## Contents

* [Installation](#installation)
* [Usage](#usage)
  * [Options](#usage-options)
  * [Limitations](#usage-limitations)
* [Generated Code](#generated)
  * [Simple](#generated-simple)
  * [Server Streaming](#generated-server)
  * [Client Streaming](#generated-client)
  * [Bidirectional Streaming](#generated-bidirectional)
* [Custom Templates](#template)
  * [Template Payload](#template-payload)
  * [Type Metadata Payload](#template-type-metadata)
  * [Method Payload](#template-method-payload)
  * [Method Type Payload](#template-type-payload)
* [Roadmap](#roadmap)

<a id="installation"></a>

## Installation

To install gServe, run:

```sh
go get github.com/Nais777/gserve
go install github.com/Nais777/gserve/...
```

gServe contains the command `protoc-gen-server` that will generate boilerplate go code for implementing gRPC servers. gServe will eventually contain a command to generate boilerplate client code also (see [Roadmap](#roadmap)).

<a id="usage"></a>

## Usage

Generation of code is done using the `protoc` command. The target destionation is defined by using the `--server_out` option. For a proto file `example.proto`, an example `protoc` execution would look like:

```sh
protoc \
    -I=. \
    -I=$GOPATH/src \
    --server_out="../grpc" \
    example.proto
```

<a id="usage-options"></a>

### Options

`package`: defines the go package for the generated files. If not specified, defaults to `grpc`.

```
  protoc \
    -I=. \
    -I=$GOPATH/src \
    --server_out="package=generated:../generated" \
    example.proto
```

`tempate`: path to a custom template used for generation of files. See [custom templates](#template) for additional information

```
  protoc \
    -I=. \
    -I=$GOPATH/src \
    --server_out="template=./test.template:../generated" \
    example.proto
```

<a id="usage-limitations"></a>

### Limitations

* Generated code cannot go into the same folder as the generated go protobuf code.

<a id="generated"></a>

## Generated Code

gServe generates one file per service in your proto definitions.

<a id="template"></a>

## Custom Templates

The default template will not suffice for everyone, so you do have the option to create your own template and supply the path to the `protoc` command. The template will be parsed using the `text/template` go library which documentation can be found [here](https://golang.org/pkg/text/template/).

The template will be executed once per service.

<a id="template-payload"></a>

### Template Payload

```go
type templatePayload struct {
    Imports      []string
    Package      string
    ProtoPackage string
    Name         string
    Methods      []methodPayload
    TypeMetadata []typeMetadata
    First        bool
}
```

`Imports` only contains the golang imports needed for the proto definitions. If you need to import any other packages, those must be defined in the template.

`Package` is the golang package for the file.

`ProtoPackage` is the protobuf package declared in the proto file.

`Name` in the name of the service being generated.

`Methods` contains information about the methods on the service.

`TypeMetadata` contains all of the types used by the service, either in requests or responses.

`First` indicates if this is the first time the template is being generated for the execution of the `protoc` command. This can be useful when you are generating code for multiple services and need to generate some code only once per execution.

<a id="template-type-metadata"></a>

### Type Metadata Payload

```go
type typeMetadata struct {
    methodType

    ClientStreamed bool
    ServerStreamed bool
}
```

See the `methodType` [section](#template-method-type) for additional properties.

`ClientStreamed` is true if the type is streamed from the client for any method.

`ServerStreamed` is true if the type is streamed from the server for any method.

<a id="template-method-payload"></a>

### Method Payload

```go
type methodPayload struct {
    Name            string
    Input           methodType
    Output          methodType
    ClientStreaming bool
    ServerStreaming bool
}
```

`Name` is the name of the method.

`Input` contains information about the input of the method.

`Output` contains information about the output of the method.

`ClientStreamed` is true if the method is a client streaming endpoint.

`ServerStreamed` is true if the method is a server streaming endpoint.

<a id="template-type-payload"></a>

### Method Type Payload

```go
type methodType struct {
    Pkg  string
    Name string
}
```

`Pkg` is the go package the type is coming from.

`Name` is the name of the type.

<a id="roadmap"></a>

## Roadmap

* Generation of bidirectional streaming endpoint code.
* Ability to supply multiple template files, one for services and one for common code.
* Ability to generate gServe code in same folder as `protoc-gen-go` code.