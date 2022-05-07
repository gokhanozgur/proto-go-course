package main

import (
	"fmt"

	pb "github.com/gokhanozgur/proto-go-course/proto"
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

func main() {
	//fmt.Println(doSimple())
	fmt.Println(doComplex())
}
