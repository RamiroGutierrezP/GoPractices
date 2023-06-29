package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	wg := sync.WaitGroup{}

	//Declaro las variables
	var numero = 1000
	var cantidad = 3

	wg.Add(cantidad)
	for i := 0; i < cantidad; i++ {
		desde := (numero / cantidad) * i + 1
		hasta := (numero / cantidad) * (i + 1) 
		go obtenerPrimos(desde, hasta, &wg)
	}
	wg.Wait()

	elapsed := time.Since(start)
    fmt.Println("Duración:", elapsed)
}

func obtenerPrimos(desde, hasta int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := desde; i <= hasta; i++ {
		if esPrimo(i) {
			fmt.Println(i)
		}
	}
}
func esPrimo(numero int) bool {
	if numero <= 1 {
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
