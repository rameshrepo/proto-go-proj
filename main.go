package main

import (
	"fmt"

	pb "github.com/rameshrepo/proto-go-proj/proto"
)

func doSimple() *pb.Simple {
	simple := pb.Simple{
		Id:          42,
		Name:        "My name",
		IsSimple:    true,
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}

	fmt.Println(simple.Id)
	fmt.Println(&simple)
	//fmt.Println(*(&simple))

	return &pb.Simple{
		Id:          42,
		Name:        "My name",
		IsSimple:    true,
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doImportOne() *pb.Import {
	return &pb.Import{One: &pb.MessageOne{Name: "My name"}}
}

func doMessageTwo() *pb.MessageTwo {
	// To understand the usage of import and package in a proto file
	return &pb.MessageTwo{Name: "MessageTwo", One: &pb.MessageOne{Name: "Message One"}}
}

func main() {
	fmt.Println(doSimple())
	// fmt.Println(doImportOne())
	//fmt.Println(doMessageTwo())
}
