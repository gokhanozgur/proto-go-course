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

func main() {
	fmt.Println(doSimple())
}
