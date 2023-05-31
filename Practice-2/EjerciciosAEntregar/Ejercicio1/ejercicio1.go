package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
	dni              string
	nombre           string
	apellido         string
	codigoDeImpuesto string
	montoAPagar      float64
}
type Codigo int
const (
	A Codigo = iota
	B
	C
	D
)
func main() {
	//Declaro las variables que voy a usar
	var clientesPorCodigo [4]int
	clientesAtendidos := 0
	montoRecaudado := 0.0
	clientes := cargarClientes()
	//clientes := make([]Cliente, 0)
	//cargarClientes(&clientes)

	// Recorro el slice de clientes
	for i:= 0; i < len(clientes) && montoRecaudado <= 10000; i++ {

		cliente := clientes[i]
		clientesAtendidos++
		montoRecaudado += cliente.montoAPagar
		incrementarContador(cliente.codigoDeImpuesto, &clientesPorCodigo)

	}
	// Imprimo los resultados
	fmt.Println("Cantidad de clientes sin atender:", len(clientes) - clientesAtendidos)
	imprimirMaximo(clientesPorCodigo)
}

func incrementarContador(codigo string, clientesPorCodigo *[4]int) {
	switch codigo {
	case "A":
		clientesPorCodigo[A]++
	case "B":
		clientesPorCodigo[B]++
	case "C":
		clientesPorCodigo[C]++
	case "D":
		clientesPorCodigo[D]++
	}
}
func imprimirMaximo(clientesPorCodigo [4]int) {
	maximo := 0
	codigoMaximo := 0
	for i := 0; i < len(clientesPorCodigo); i++ {
		if clientesPorCodigo[i] > maximo {
			maximo = clientesPorCodigo[i]
			codigoMaximo = i
		}
	}
	switch codigoMaximo {
	case 0:
		fmt.Println("El código con más clientes es A")
	case 1:
		fmt.Println("El código con más clientes es B")
	case 2:
		fmt.Println("El código con más clientes es C")
	case 3:
		fmt.Println("El código con más clientes es D")
	}
}
func convertirStringAFloat(numero string) float64 {
	valor, err := strconv.ParseFloat(numero, 64)
	if err != nil {
		log.Fatal("Error al convertir el número: ", numero)
	}
	return valor
}

func cargarClientes() []Cliente {

	clientes := make([]Cliente, 0)
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		var cliente Cliente
		
		//Leo la linea y separo los datos por campo
		linea := scanner.Text()
		datos := strings.Split(linea, ",")

		//Valido que la linea tenga 5 campos; si no tiene salteo la linea
		if len(datos) != 5 {
			continue
		}

		cliente.dni = datos[0]
		cliente.nombre = datos[1]
		cliente.apellido = datos[2]
		cliente.codigoDeImpuesto = datos[3]
		cliente.montoAPagar = convertirStringAFloat(datos[4])

		clientes = append(clientes, cliente)
	}
	return clientes
}

// func cargarClientes(clientes *[]Cliente) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		var cliente Cliente
// 		//Leo la linea y separo los datos por campo
// 		linea := scanner.Text()
// 		datos := strings.Split(linea, ",")
//
// 		//Valido que la linea tenga 5 campos; si no tiene salteo la linea
// 		if len(datos) != 5 {
// 			continue
// 		}
//
// 		cliente.dni = datos[0]
// 		cliente.nombre = datos[1]
// 		cliente.apellido = datos[2]
// 		cliente.codigoDeImpuesto = datos[3]
// 		cliente.montoAPagar = convertirStringAFloat(datos[4])
//
// 		*clientes = append(*clientes, cliente)
// 	}
// }

