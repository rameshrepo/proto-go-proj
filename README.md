go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc -Iproto --go_opt=module=github.com/rameshrepo/proto-go-proj --go_out=. proto/*.proto