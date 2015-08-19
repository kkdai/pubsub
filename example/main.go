package main

import (
	"fmt"

	. "../../go-pubsub"
)

func main() {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("topic1")
	ser.Publish("test1", "topic1")
	fmt.Println(<-c1)
}
