package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	pb "proto-go-course/internal/proto"
	"reflect"
)

func main() {
	//jsonString := doToJSON(doSimple())
	//message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))

	jsonString := doToJSON(doComplex())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))

	fmt.Println(jsonString)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	fmt.Println(jsonString)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}

	readFromFile(path, message)
	fmt.Println(message)
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

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"Nem": {
				Id: 43,
			},
			"Ryan": {
				Id: 56,
			},
			"Pete": {
				Id: 9,
			},
		},
	}
}
