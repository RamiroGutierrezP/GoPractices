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
	clientes := make([]Cliente, 0)
	cantidadClientes := 0
	clientesAtendidos := 0
	montoTotal := 0.0
	var clientesPorCodigo [4]int
	scanner := bufio.NewScanner(os.Stdin)

	// Recorro el archivo para leer todos los clientes
	for scanner.Scan() {
		
		// Creo una variable cliente para guardar los datos de cada cliente
		var cliente Cliente

		// Leo una linea del archivo
		linea := scanner.Text()
		datos := strings.Split(linea, ",")

		// Valido que la linea tenga 5 datos; si no tiene salteo la linea
		if len(datos) != 5 {
			log.Print("Formato de cliente inválido en línea: ", linea)
			continue
		}

		// Cargo los datos del cliente
		cliente.dni = datos[0]
		cliente.nombre = datos[1]
		cliente.apellido = datos[2]
		cliente.codigoDeImpuesto = datos[3]
		cliente.montoAPagar, _ = strconv.ParseFloat(datos[4], 64)

		// Agrego el cliente al slice de clientes
		clientes = append(clientes, cliente)
	
		//Hago los cálculos solicitados
		if montoTotal < 10000 {
			clientesAtendidos++
			montoTotal += cliente.montoAPagar
			incrementarContador(cliente.codigoDeImpuesto, &clientesPorCodigo)
		}
		cantidadClientes++
	}

	// Imprimo los resultados
	fmt.Println("Cantidad de clientes sin atender:", cantidadClientes - clientesAtendidos)
	// fmt.Println("Monto total:", montoTotal)
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