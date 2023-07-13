package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func main() {
	start := time.Now()
	
	cantidadCajas := 3
	cantidadClientes := 3
	
	var wg sync.WaitGroup
	colas := crearColasConBuffer(cantidadCajas)
	cajas := crearCanales(cantidadCajas)
	go clientesLlegando(colas, cantidadClientes, cantidadCajas)
	wg.Add(cantidadCajas)
	for i := 0; i < cantidadCajas; i++ {
		go func(i int) {
			for cliente := range colas[i] {
				go atenderCliente(cliente, i, cajas[i])
				<-cajas[i]
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	
	elapsed := time.Since(start)
    fmt.Println("Duración:", elapsed)
}
func crearColasConBuffer(cantidad int) []chan int{
	colas := make([]chan int, cantidad)
	for i := 0; i < cantidad; i++ {
		colas[i] = make(chan int, 100)
	}
	return colas
}
func clientesLlegando(colas []chan int, cantidadClientes int, cantidadCajas int) {
	for i := 0; i < cantidadClientes; i++ {
		//Envío el cliente a la cola que menos clientes tenga
		obtenerColaConMenosClientes(colas) <- i		
	}
	for i := 0; i < cantidadCajas; i++ {
		close(colas[i])
	}
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