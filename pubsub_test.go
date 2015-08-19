package pubsub

import (
	"fmt"
	"testing"
)

func TestBasicFunction(t *testing.T) {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("ch1")
	ser.Publish("test1", "ch1")

	if val, ok := <-c1; ok {
		fmt.Printf(" Got content from subscribed topic %v\n", val)
	} else {
		t.Error(" Error found on subscribed.\n")
	}
}

func TestTwoSubscribetor(t *testing.T) {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("ch1")
	c2 := ser.Subscribe("ch2")

	ser.Publish("test2", "ch1")
	ser.Publish("test1", "ch2")

	if val, ok := <-c1; ok && val == "test2" {
		fmt.Printf("ret: %v \n", val)
	} else {
		t.Errorf("Error found \n")
	}

	fmt.Printf("c2= %v \n", <-c2)
}

func TestAddSub(t *testing.T) {
	ser := NewPubsub(10)
	c1 := ser.Subscribe("ch1")
	c2 := ser.Subscribe("ch1")
	ser.AddSubscription(c1, "ch2")
	ser.Publish("test2", "ch1")

	fmt.Printf("first c1:%v \n", <-c1)
	fmt.Printf("c2= %v \n", <-c2)
}
