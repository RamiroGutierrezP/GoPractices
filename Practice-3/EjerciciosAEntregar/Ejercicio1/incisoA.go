package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	var numero = 100000

	wg.Add(1)
	go obtenerPrimos(numero, &wg)
	wg.Wait()

	elapsed := time.Since(start)

    fmt.Println("Speed-up:", elapsed)
}

func obtenerPrimos(numero int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= numero; i++ {
		if esPrimo(i) {
			fmt.Println(i)
		}
	}
}
func esPrimo(numero int) bool {
	if numero == 1 {
		return false
	}
	//Verifica si es divisible por cada numero menor a si mismo
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}