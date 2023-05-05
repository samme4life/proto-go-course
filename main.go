package main

import (
	"fmt"

	pb "github.com/samme4life/proto-go-course/proto"
)

func main() {
	fmt.Println(doComplex())
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "Alex",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id: 23,
			Name: "OneDummy",
		},
		MultipleDummies: []*pb.Dummy{
			&pb.Dummy{
				Id: 41,
				Name: "TwoDummy",
			},
			&pb.Dummy{
				Id: 23,
				Name: "ThreeDummy",
			},
		},
	}
}
