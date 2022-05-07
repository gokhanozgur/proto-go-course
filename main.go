package main

import (
	"fmt"
	"reflect"

	pb "github.com/gokhanozgur/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         28,
		IsSimple:   true,
		Name:       "Gokhan",
		SampleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   28,
			Name: "Gokhan",
		},
		MultipleDummies: []*pb.Dummy{
			{Id: 21, Name: "Gizem"},
			{Id: 53, Name: "Yucel"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		//EyeColor: 1,
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 28},
			"myid2": {Id: 21},
			"myid3": {Id: 53},
		},
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type %v", x)
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJson(p)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJson(jsonString, message)
	return message
}

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//fmt.Println(doEnum())
	//fmt.Println(doMap())
	/*
		fmt.Print("This should be an Id: ")
		doOneOf(&pb.Result_Id{Id: 28})
		fmt.Print("This should be an message: ")
		doOneOf(&pb.Result_Message{Message: "a message"})
	*/
	//doFile(doSimple())
	/*
		jsonString := doToJSON(doSimple())
		message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
		fmt.Println(jsonString)
		fmt.Println(message)
	*/

	/*
		jsonString = doToJSON(doComplex())
		message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
		fmt.Println(jsonString)
		fmt.Println(message)
	*/

	fmt.Println(doFromJSON(`{"id": 28, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))
}
