## Protobuf for Golang

### Getting Started
1. Install `protoc`
```
$ sudo apt-get update && sudo apt-get install -y protobuf-compiler 
```

2. Install proto plugin for golang. 
The above plugin is needed to create go files for proto using protoc
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
protoc-gen-go needs to be in your shell path, i.e. one of the directories listed in the PATH environment variable, which is different from the Go path. You can test this by simply typing protoc-gen-go at the command line: If it says "command not found" (or similar) then it's not in your PATH. 

Check the location of protoc-gen-go. it should live at \$GOPATH/bin. So do this:
```bash
$ export PATH=$PATH:$GOPATH/bin 
(or)
$ export PATH=$PATH:$(go env GOPATH)/bin
```
3. Generate proto files `protoc`
```
$ protoc -Iproto --go_opt=module=github.com/rameshrepo/proto-go-proj --go_out=. proto/*.proto
```
4. Development inside Docker

Instead of installing golang, protobuf compiler and protoc plugins for go by executing the above 3 steps in your local. You can instead spin up the development setup inside Docker.
```
docker build --build-arg GO_VERSION=1.20.6 --tag proto-go-proj . 
docker run -d -v ${PWD}:/apps/proto-go-proj --name proto-go-proj -it proto-go-proj
```
To Mount the current directory to /apps/proto-go-proj inside docker
<dl>
  <dt>Powershell</dt>
  <dd>docker run -d -v ${PWD}:/apps/proto-go-proj --name proto-go-proj -it proto-go-proj</dd>
  <dt>Windows</dt>
  <dd>docker run -d -v %cd%:/apps/proto-go-proj --name proto-go-proj -it proto-go-proj</dd>
  <dt>Linux</dt>
  <dd>docker run -d -v $(pwd):/apps/proto-go-proj --name proto-go-proj -it proto-go-proj</dd>
</dl>

### Protobuf Basics

The compiler generates a *.pb.go file for each .proto file.