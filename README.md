Pubsub
==============

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/pubsub/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/pubsub?status.svg)](https://godoc.org/github.com/kkdai/pubsub)  [![Build Status](https://travis-ci.org/kkdai/pubsub.svg?branch=master)](https://travis-ci.org/kkdai/pubsub)



What is Pubsub
=============
Pubsub is prove of concept implement for [Redis](http://redis.io/) "Pub/Sub" messaging management feature. SUBSCRIBE, UNSUBSCRIBE and PUBLISH implement the Publish/Subscribe messaging paradigm where (citing Wikipedia) senders (publishers) are not programmed to send their messages to specific receivers (subscribers).  (sited from [here](http://redis.io/topics/pubsub))



Install
---------------
`go get github.com/kkdai/pubsub`


Usage
---------------

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


Inspired By
=============


- [Redis: Pubsub](http://redis.io/topics/pubsub)
- [chandru/pubsub](https://github.com/tuxychandru/pubsub)


License
---------------

This package is licensed under MIT license. See LICENSE for details.
