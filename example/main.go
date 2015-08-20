package main

import (
	"fmt"

	. "github.com/kkdai/pubsub"
)

func main() {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("topic1")
	c2 := ser.Subscribe("topic2")
	ser.Publish("test1", "topic1")
	ser.Publish("test2", "topic2")
	fmt.Println(<-c1)
	//Got "test1"
	fmt.Println(<-c2)
	//Got "test2"
}
