package pubsub_test

import (
	"fmt"

	. "github.com/kkdai/pubsub"
)

func ExamplePubsub() {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("topic1")
	c2 := ser.Subscribe("topic2")

	//Add subscription for c1
	ser.AddSubscription(c1, "topic2", "topic3")
	//Remove subscription topic2 in c1
	ser.RemoveSubscription(c1, "topic2")

	//Publish content "test1" in topic "topic1"
	ser.Publish("test1", "topic1")
	//Publish content "test2" in topic "topic2"
	ser.Publish("test2", "topic2")

	fmt.Println(<-c1)
	//Got "test1"
	fmt.Println(<-c2)
	//Got "test2"
}

func ExamplePubsub_Subscribe() {
	ser := NewPubsub(1)
	//Subscribe topic1 in c1
	c1 := ser.Subscribe("topic1")
	fmt.Printf("c1 is a channel %v\n", c1)
	//Subscribe multiple topic "topic1", "topic2" and "topic3" in c2
	c2 := ser.Subscribe("topic1", "topic2", "topic3")
	fmt.Printf("c2 is a channel %v\n", c2)
}

func ExamplePubsub_AddSubscription() {
	ser := NewPubsub(1)
	//Subscribe topic1 in c1
	c1 := ser.Subscribe("topic1")
	//Add more Subscription "topic2" and "topic3" for c1
	ser.AddSubscription(c1, "topic2", "topic3")
}
