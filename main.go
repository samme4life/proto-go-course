package main

import (
	"fmt"

	pb "github.com/samme4life/proto-go-course/proto"
)

func main() {
	fmt.Println(doSimple())
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "Alex",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}
