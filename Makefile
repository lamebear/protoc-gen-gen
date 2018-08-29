build-example:
	protoc -I=./example \
		-I=${GOPATH}/src/github.com/google/protobuf/src \
		--go_out="plugins=grpc:${GOPATH}/src" \
		--gen_out="package=example,template=./example/service.template:./example/generated" \
		./example/example.proto