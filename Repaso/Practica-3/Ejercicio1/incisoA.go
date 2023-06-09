package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	
	done := make(chan bool)
	// var wg sync.WaitGroup

	var numero = 100000

	// wg.Add(1)
	go obtenerPrimos(numero, done)
	// wg.Wait()
	<-done

	elapsed := time.Since(start)
    fmt.Println("Duración:", elapsed)
}

func obtenerPrimos(numero int, done chan bool) {
	// defer wg.Done()
	for i := 1; i <= numero; i++ {
		if esPrimo(i) {
			fmt.Println(i)
		}
	}
	done <- true
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