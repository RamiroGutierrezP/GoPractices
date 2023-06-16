package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Inicia Goroutine del main")
	done := make(chan bool)
	go hello(done)
	<- done
	fmt.Println("Termina Goroutine del main")
}
func hello(done chan bool) {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Termina Goroutine de hello")
	done <- true
}
