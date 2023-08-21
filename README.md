## Learning Protobuf for Golang

### Getting Started
1. Install `protoc`
```
protoc -Iproto --go_opt=module=github.com/rameshrepo/proto-go-proj --go_out=. proto/*.proto
```
2. Install proto plugin for golang
The above plugin is needed to create a go files for proto using protoc
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
protoc-gen-go needs to be in your shell path, i.e. one of the directories listed in the PATH environment variable, which is different from the Go path. You can test this by simply typing protoc-gen-go at the command line: If it says "command not found" (or similar) then it's not in your PATH. 

Check the location of protoc-gen-go. it should live at \$GOPATH/bin. So do this:
\$ export PATH=\$PATH:\$GOPATH/bin 
(or)
\$ export PATH="\$PATH:$(go env GOPATH)/bin"

### Protobuf Basics

The compiler generates a *.pb.go file for each .proto file.