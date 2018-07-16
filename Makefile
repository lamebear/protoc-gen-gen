build-example:
	protoc -I=./example \
		-I=${GOPATH}/src/github.com/google/protobuf/src \
		--go_out=plugins=grpc:${GOPATH}/src \
		--server_out=package=example:./example/generated \
		./example/example.proto

build-template:
	go generate ./template/...