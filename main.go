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

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//fmt.Println(doEnum())
	fmt.Println(doMap())
}
