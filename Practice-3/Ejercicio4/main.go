package main

import (
	"fmt"
	"sync"
)
func pxng(m chan string, str string, wg *sync.WaitGroup) {
	defer wg.Done()
	m <- str
}

func main() {
	var wg sync.WaitGroup
	messages := make(chan string, 10)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go pxng(messages, "PING", &wg)
		wg.Wait()
		wg.Add(1)
		go pxng(messages, "PONG", &wg)
		wg.Wait()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-messages)
	}
}