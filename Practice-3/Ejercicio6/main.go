package main

import (
	"fmt"
	"math/rand"
)

func main() {

	one := make(chan int, 5)
	two := make(chan int, 5)
	three := make(chan int, 5)

	go fillChannel(one)
	go fillChannel(two)
	go fillChannel(three)


	for i:= 0; i < 15; i++{
		select {
		case number := <-one:
			fmt.Println("Channel 1: ", number)
		case number := <-two:
			fmt.Println("Channel 2: ", number)
		case number := <-three:
			fmt.Println("Channel 3: ", number)
		}
	}
	
}

func fillChannel(channel chan int) {
	// defer wg.Done()
	for i := 0; i < 5; i++ {
		channel <- rand.Intn(100) + 1
	}
}
