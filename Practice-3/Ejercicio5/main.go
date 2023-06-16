package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	//Channel for the numbers 
	numbers := make(chan int)

	//Productors
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go productor(numbers, &wg, i)
	}
	//Consumers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go consumer(numbers, &wg, i)
	}
	wg.Wait()
}

func productor(numbers chan<- int, wg *sync.WaitGroup, id int){
	
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(1)) * time.Second)

	for i := 0; i < 3; i++ {
		numbers <- rand.Intn(100) + 1
	}
}
func consumer (numbers <-chan int, wg *sync.WaitGroup, id int){
	defer wg.Done()

	for i := 0; i < 3; i++ {
		number := <- numbers
		fmt.Printf("Consumer %d: %d\n", id, number)
	}
}


