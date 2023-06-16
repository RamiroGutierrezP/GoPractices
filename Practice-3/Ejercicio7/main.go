package main

import (
	"math/rand"
	"time"
)

func main() {

	//Channels
	one := make(chan int)
	two := make(chan int)

	go sender1(one)
	go sender2(two)

	timeout1 := time.After(5 * time.Second)
	timeout2 := time.After(10 * time.Second)

	for {
		select {
		case number, ok := <-one:
			if !ok {
				one = nil
				break
			}
			println("Channel 1: ", number)
		case number, ok  := <-two:
			if !ok {
				two = nil
				break
			}
			println("Channel 2: ", number)
		case <-timeout1:
			println("Timeout 1")
			one = nil
		case <-timeout2:
			println("Timeout 2")
			two = nil
		}
		if one == nil && two == nil {
			break
		}
	}

}

func sender1 (channel chan int) {
	for {
		channel <- rand.Intn(100) + 1
		time.Sleep(1 * time.Second)
	}
}
func sender2 (channel chan int) {
	for {
		channel <- rand.Intn(100) + 1
		time.Sleep(1 * time.Second)
	}
}