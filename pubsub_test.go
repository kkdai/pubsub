package pubsub

import (
	"testing"
	"time"
)

func TestBasicFunction(t *testing.T) {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("ch1")
	ser.Publish("test1", "ch1")

	if _, ok := <-c1; !ok {
		t.Error(" Error found on subscribed.\n")
	}
}

func TestTwoSubscribetor(t *testing.T) {
	ser := NewPubsub(1)
	c1 := ser.Subscribe("ch1")
	c2 := ser.Subscribe("ch2")

	ser.Publish("test2", "ch1")
	ser.Publish("test1", "ch2")

	val, ok := <-c1
	if !ok || val != "test2" {
		t.Errorf("Error found \n")
	}

	val, ok = <-c2
	if !ok || val != "test1" {
		t.Errorf("Error found \n")
	}
}

func TestAddSub(t *testing.T) {
	ser := NewPubsub(10)
	c1 := ser.Subscribe("ch1")
	ser.AddSubscription(c1, "ch2")
	ser.Publish("test2", "ch2")

	if val, ok := <-c1; !ok {
		//Not get! Occur error.
		t.Errorf("error on c1:%v", val)
	}
}

func TestRemoveSub(t *testing.T) {
	ser := NewPubsub(10)
	c1 := ser.Subscribe("ch1", "ch2")
	ser.Publish("test1", "ch2")

	if val, ok := <-c1; !ok {
		t.Errorf("error on addsub c1:%v", val)
	}

	ser.RemoveSubscription(c1, "ch1")
	ser.Publish("test2", "ch1")

	select {
	case val := <-c1:
		t.Errorf("Should not get %v notify on remove topic\n", val)
		break
	case <-time.After(time.Second):
		break
	}
	//if val, ok := <-c1; ok {
	//	t.Errorf("error on remove Sub c1:", val)
	//}

}
