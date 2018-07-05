package main

import (
	"fmt"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func fanIn2(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Second * 1)
			<-waitForIt
		}
	}()
	return c
}

func fan2() {
	c := fanIn2(boring("Joe"), boring("Ann"))

	for i := 0; i <= 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}
