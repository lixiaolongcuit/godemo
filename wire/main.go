package main

import "fmt"

type Message struct {
	msg string
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

//Message构造函数
func NewMessage(msg string) Message {
	return Message{msg: msg}
}

//Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

//Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// 使用wire前
// func main() {
// 	message := NewMessage("hello world")
// 	greeter := NewGreeter(message)
// 	event := NewEvent(greeter)

// 	event.Start()
// }

//使用wire后
// 一起编译才能运行
// go build main.go wire_gen.go
func main() {
	event := InitializeEvent("hello wire")

	event.Start()
}
