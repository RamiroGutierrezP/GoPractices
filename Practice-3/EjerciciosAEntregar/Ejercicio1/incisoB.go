package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	//Declaro las variables
	var numero = 100000
	primeraMitad := make(chan int)
	segundaMitad := make(chan int)

	//Inicio las goroutines
	go obtenerPrimosPrimeraMitad(numero, primeraMitad)
	go obtenerPrimosSegundaMitad(numero, segundaMitad)
	
	//Mientras las goroutines no terminen, imprimo los numeros que vayan llegando
	stop1, stop2 := false, false
	for !stop1 || !stop2 {
		select {
		case numero, ok := <-primeraMitad:
			if ok {
				fmt.Println(numero)
			} else {
				stop1 = true
			}
		case numero, ok2 := <-segundaMitad:
			if ok2 {
				fmt.Println(numero)
			} else {
				stop2 = true
			}
		}
	}

	elapsed := time.Since(start)
    fmt.Println("Speed-up:", elapsed)
}

func obtenerPrimosPrimeraMitad(numero int, primos chan<- int) {
	for i := 1; i < numero/2; i++ {
		if esPrimo(i) {
			primos <- i
		}
	}
	close(primos)
}
func obtenerPrimosSegundaMitad(numero int, primos chan<- int) {
	for i := numero/2; i <= numero; i++ {
		if esPrimo(i) {
			primos <- i
		}
	}
	close(primos)
}
func esPrimo(numero int) bool {
	//Podria sacarlo si inicializo el for en 2
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
