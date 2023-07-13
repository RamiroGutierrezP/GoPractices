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

func main() {
	//Declaro las variables que voy a usar
	clientesPorCodigo := make(map[string]int)
	clientesAtendidos := 0
	montoRecaudado := 0.0
	clientes := cargarClientes()

	// Recorro el slice de clientes
	for i:= 0; i < len(clientes) && montoRecaudado <= 10000; i++ {

		cliente := clientes[i]
		clientesAtendidos++
		montoRecaudado += cliente.montoAPagar
		clientesPorCodigo[cliente.codigoDeImpuesto]++

	}
	// Imprimo los resultados
	fmt.Println("Cantidad de clientes sin atender:", len(clientes) - clientesAtendidos)
	imprimirMaximo(clientesPorCodigo)
}

func imprimirMaximo(clientesPorCodigo map[string]int) {
	maximo := 0
	codigoMaximo := ""
	for codigo, cantidad := range clientesPorCodigo {
		if cantidad > maximo {
			maximo = cantidad
			codigoMaximo = codigo
		}
	}
	fmt.Println("El código con más clientes es", codigoMaximo)
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

