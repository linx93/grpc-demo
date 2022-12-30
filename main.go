package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"grpc-demo/service"
)

func main() {
	password := "123456"

	stu := &service.Student{
		Username: "linx",
		Password: &password,
		Age:      18,
		Gender:   1,
		Class: &service.Class{
			ClassName: "一年级一班",
			ClassCode: "1-01",
		},
		Addresses: []string{"北京天安门", "南京"},
		Teachers:  []*service.Teacher{{Age: 18, Gender: 1, TeacherCode: "007", TeacherName: "的马西亚"}},
	}

	//序列化
	marshal, err := proto.Marshal(stu)
	if err != nil {
		panic(err)
	}

	newStu := &service.Student{}
	err = proto.Unmarshal(marshal, newStu)
	if err != nil {
		panic(err)
	}

	fmt.Println(newStu)

}
