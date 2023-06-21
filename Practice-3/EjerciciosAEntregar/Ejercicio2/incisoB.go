package main

import (
	"math/rand"
	"sync"
	"time"
)


func main() {
	wg := sync.WaitGroup{}
	colas := crearCanales(3)
	cajas := crearCanales(3)
	go clientesLlegando(colas)
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for cliente := range colas[i] {
				go atenderCliente(cliente, i, cajas[i])
				<-cajas[i]
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func clientesLlegando(colas []chan int) {
	for i := 0; i < 22; i++ {
		//Envío el cliente a la cola correspondiente mediante el algoritmo de round robin
		colas[i % 3] <- i
	}
	for i := 0; i < 3; i++ {
		close(colas[i])
	}
}
func crearCanales(cantidad int) []chan int{
	canales := make([]chan int, cantidad)
	for i := 0; i < cantidad; i++ {
		canales[i] = make(chan int)
	}
	return canales
}
func atenderCliente(cliente int, caja int, cajaCanal chan int) {

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	cajaCanal <- cliente
	println("Atendiendo al cliente: ", cliente+1, " en la caja: ", caja+1)
}