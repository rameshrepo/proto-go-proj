package main

import (
	"fmt"

	pb "github.com/rameshrepo/proto-go-proj/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		Name:        "My name",
		IsSimple:    true,
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doMessageTwo() *pb.MessageTwo {
	// To understand the usage of import and package in a proto file
	return &pb.MessageTwo{Name: "MessageTwo", One: &pb.MessageOne{Name: "Message One"}}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doMapObjExample() *pb.MapObjExample {
	message := &pb.MapObjExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
	return message
}

func doMapExample() *pb.MapExample {
	message := &pb.MapExample{
		Citypostcodes: map[string]uint32{
			"kingsville": 3012,
			"pointcook":  3030,
		},
	}
	return message
}

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doMessageTwo())
	//fmt.Println(doEnum())
	fmt.Println(doMapObjExample())
	fmt.Println(doMapExample())
}
