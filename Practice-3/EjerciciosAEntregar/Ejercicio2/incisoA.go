package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	
	caja := make(chan int)
	cola := make(chan int)
	go clientesLlegando(cola)

	for cliente := range cola {
		go atenderCliente(cliente, caja)
		<-caja
	}
	close(caja)

	elapsed := time.Since(start)
    fmt.Println("Speed-up:", elapsed)
}

//Simulo la llegada de clientes
func clientesLlegando(cola chan int) {
	for i := 1; i <= 10; i++ {
		cola <- i
	}
	close(cola)
}
//Simulo la atención del cliente
func atenderCliente(cliente int, caja chan int){
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	caja <- cliente
	println("Atendiendo cliente", cliente)
}
