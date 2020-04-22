package main

import (
	"godemo/protobuf/msg"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	h := &msg.Hello{
		Name: "jerry",
	}
	//序列号
	data, err := proto.Marshal(h)
	if err != nil {
		log.Fatal("pb Marshal error:", err)
	}
	log.Println("Data:%v", data)
	h2 := &msg.Hello{}
	err = proto.Unmarshal(data, h2)
	if err != nil {
		log.Fatal("pb Unmarshal error:", err)
	}
	log.Println("hello:", h2)
}
