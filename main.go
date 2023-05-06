package main

import (
	"fmt"
	pb "proto-go-course/internal/proto"
)

func main() {
	doOneOf(&pb.Result_Id{
		Id: 76,
	})
	doOneOf(&pb.Result_Message{
		Message: "Hello World!",
	})
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
			Id:   23,
			Name: "OneDummy",
		},
		MultipleDummies: []*pb.Dummy{
			&pb.Dummy{
				Id:   41,
				Name: "TwoDummy",
			},
			&pb.Dummy{
				Id:   23,
				Name: "ThreeDummy",
			},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColour: pb.EyeColour_EYE_COLOUR_BLUE,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}
}
