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

To install gServe, simply run:

```sh
go get github.com/Nais777/gserve
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

`tempate`: path to a custom template used for generation of files. See [custom templates](#payload) for additional information

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

An `interface` type will be defined that will need to be injected into the generated struct.

When an error is returned from the injected struct, that error should implement the following interface when possible:

```go
type grpcError interface {
    Code() codes.Code
}
```

The `Code()` function will be used to send the appriate grpc error code to the caller. If the error does not implement this interface, the error code `codes.Unknown` will be sent.

For the following sections, we will be using the following proto file:

```proto3
syntax="proto3";

option go_package="github.com/Nais777/gserve/example";

package example;

import "google/protobuf/empty.proto";

message ID {
    int32 ID = 1;
}

message User {
    example.ID ID = 1;
    string FirstName = 2;
    string LastName = 3;
}

service Users {
    rpc User(example.ID) returns (example.User) {}
    rpc Users(google.protobuf.Empty) returns (stream example.User) {}
    rpc BatchUpdate(stream example.User) returns (stream example.User) {}
    rpc BatchDelete(stream example.ID) returns (google.protobuf.Empty) {}
}
```

<a id="generated-simple"></a>

### Simple RPC

As there is no special processing that needs to occur for a simple rpc, the generated code simply passes the request to the injected service.

Protobuf Definition:

```protobuf
service Users {
    rpc User(example.ID) returns (example.User) {}
}
```

Generated Interface Func:

```go
type UsersService interface {
    User(ctx context.Context, req *example.ID) (*example.User, error)
}
```

<a id="generated-server"></a>

### Server Streaming RPC

The injected service returns a cursor object to the generated server function that will be used to iterate over the data to be returned. In addition, the cursor also has an `Err()` function that should be used to comunicate any error that occurs in preparing the data to be sent.

By using a cursor object instead of `chan` or a `slice` allows flexibility for the implementation. The cursor struct could be backed by `slice` or a `chan` and it will be abstracted from the generated function. This also does not force the use of a `goroutine` in the case that it is not suited for the implementation.

The `Next()` method on the cursor should return `true` while there is still data to be sent, returning `false` once all the data has been iterated. `Next()` should also return `false` in the case of an error.

`Err()` on the cursor is not called until `Next()` returns false, so it is important that when an error does occur, that `Next()` returns false so the error can be sent to the client.

Protobuf Definition:

```protobuf
service Users {
    rpc Users(google.protobuf.Empty) returns (stream example.User) {}
}
```

Generated Interfaces:

```go
type UsersService interface {
    Users(ctx context.Context, req *empty.Empty) (UserCursor, error)
}

type UserCursor interface {
    Next() bool
    Current() *example.User
    Err() error
}
```

<a id="generated-client"></a>

### Client Streaming RPC

The generated server function will begin by calling the `Begin{{MethodName}}` function on the injected service. This function would be responsible for doing any required setup to handle the incoming stream of data (begining a transaction, starting `goroutine`s, etc). This function also returns a struct that will be used to process the incoming data stream

The `Next({{stream DataType}})` func will be called for every item sent to the server from the client. If there are any errors in processing the streamed data, return an error and it will be propogated to the client.

After all the data has been received from the client, `Complete()` will be called on the returned interface indicating that there is no more data to be received from the client. This function returns the struct that will be supplied to the called using the `SendAndClose` func on the gRPC stream or an error that will be propogated to the client.

`Cancel()` on the returned struct will be defered immediately following the call to `Begin{{MethodName}}`. This function will be called on every invocation. Care should be used when writing the implementing struct that calling `Cancel()` after `Complete()` returns successfully does not cause errors or a `panic`.

Protobuf Definition:

```protobuf
service Users {
    rpc BatchDelete(stream example.ID) returns (google.protobuf.Empty) {}
}
```

Generated Interfaces:

```go
type UsersService interface {
    BeginBatchDelete(ctx context.Context) (BatchDeleter, error)
}

type BatchDeleter interface {
    Next(*example.ID) error
    Complete() (*empty.Empty, error)
    Cancel()
}
```

<a id="generated-bidirectional"></a>

### Bidirectional Streaming RPC

Currently, gServe does not generate any processing code for bidirectional streaming endpoints. This is because I could not settle on a good way for handling this streaming method. It is planned, and on the [Roadmap](#roadmap).

Protobuf Definition:

```protobuf
service Users {
    rpc BatchUpdate(stream example.User) returns (stream example.User) {}
}
```

Generated Interface Func:

```go
type UsersService interface {
    BatchUpdate(stream example.Users_BatchUpdateServer) error
}
```

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
* Addition of `protoc-gen-client` to generate client code.