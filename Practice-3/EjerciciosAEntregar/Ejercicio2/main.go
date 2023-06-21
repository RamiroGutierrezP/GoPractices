package main

import (
	"math/rand"
	"sync"
	"time"
)


func main() {
	wg := sync.WaitGroup{}
	colas := crearColasConBuffer(3)
	go clientesLlegando(colas)

	for i := 0; i < 5; i++ {
		//Envío el cliente a la cola que menos clientes tenga
		obtenerColaConMenosClientes(colas) <- i		
	}
	for i := 0; i < 3; i++ {
		close(colas[i])
	}
}
func crearColasConBuffer(cantidad int) []chan int{
	colas := make([]chan int, cantidad)
	for i := 0; i < cantidad; i++ {
		colas[i] = make(chan int, 7)
	}
	return colas
}

func obtenerColaConMenosClientes(colas []chan int) chan int{
	min := len(colas[0])
	minIndex := 0
	for i := 1; i < len(colas); i++ {
		if len(colas[i]) < min {
			min = len(colas[i])
			minIndex = i
		}
	}
	return colas[minIndex]
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