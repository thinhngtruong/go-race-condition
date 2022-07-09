package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		randTime := rand.Intn(5)
		time.Sleep(time.Duration(randTime) * time.Second)
		c1 <- "one"
	}()
	go func() {
		randTime := rand.Intn(5)
		time.Sleep(time.Duration(randTime) * time.Second)
		c2 <- "two"
	}()

	var selected string
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			selected = msg1
		case msg2 := <-c2:
			selected = msg2
		}
	}

	fmt.Println("received", selected)
}
