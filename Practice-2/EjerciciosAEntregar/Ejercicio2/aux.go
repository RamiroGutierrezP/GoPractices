package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ingresante struct {
	apellido string
	nombre string
	ciudadOrigen string
	fechaNacimiento Fecha
	presentoTitulo bool
	codigoDeCarrera string
}

type Fecha struct {
	dia int
	mes int
	anio int
}

func main() {
	
	//Creo las variables
	var ingresantes []Ingresante
	var contadorDeCarreras [3]int
	var aniosDeNacimiento []int

	scanner := bufio.NewScanner(os.Stdin)

	//Recorro el archivo 
	for scanner.Scan() {
		var ingresante Ingresante

		linea := scanner.Text()
		datos := strings.Split(linea, ",")

		// Valido que la linea tenga 8 datos; si no tiene salteo la linea
		if len(datos) != 8 {
			log.Print("Formato de ingresante inválido en línea: ", linea)
			continue
		}

		// Cargo los datos 
		ingresante.apellido = datos[0]
		ingresante.nombre = datos[1]
		ingresante.ciudadOrigen = datos[2]
		ingresante.fechaNacimiento.dia = convertirStringAInt(datos[3])
		ingresante.fechaNacimiento.mes = convertirStringAInt(datos[4])
		ingresante.fechaNacimiento.anio = convertirStringAInt(datos[5])
		ingresante.presentoTitulo = convertirStringABool(datos[6])
		ingresante.codigoDeCarrera = datos[7]
		// Agrego el ingresante al slice
		ingresantes = append(ingresantes, ingresante)

		//Si es de Bariloche informo nombre y apellido
		if ingresante.ciudadOrigen == "Bariloche" {
			fmt.Println(ingresante.nombre, ingresante.apellido)
		}
		//Si no presenta titulo lo tengo que eliminar de la lista
		if ingresante.presentoTitulo == false {
			ingresantes = ingresantes[:len(ingresantes)-1]
		}
		//Cuento la cantidad de ingresantes por carrera
		contadorDeCarreras = incrementarContador(ingresante.codigoDeCarrera, contadorDeCarreras)
		// Incrementar año de nacimiento
		aniosDeNacimiento = append(aniosDeNacimiento, ingresante.fechaNacimiento.anio)
	}

	imprimirCarreraConMasIngresantes(contadorDeCarreras)
	imprimirAnioDeNacimientoConMasIngresantes(aniosDeNacimiento)

}


func convertirStringAInt(numero string) int {
	valor, err := strconv.Atoi(numero)
	if err != nil {
		fmt.Println("Error al convertir el string a int")
	}
	return valor
}
func convertirStringABool(booleano string) bool {
	valor, err := strconv.ParseBool(booleano)
	if err != nil {
		fmt.Println("Error al convertir el string a bool")
	}
	return valor
}

func incrementarContador(codigoDeCarrera string, contadorDeCarreras [3]int) [3]int {
	switch codigoDeCarrera {
	case "APU":
		contadorDeCarreras[0]++
	case "LI":
		contadorDeCarreras[1]++
	case "LS":
		contadorDeCarreras[2]++
	default:
		fmt.Println("El codigo de carrera", codigoDeCarrera ,"no es valido")
	}
	return contadorDeCarreras
}

func imprimirCarreraConMasIngresantes (contadorDeCarreras [3]int) {
	var carreraConMasIngresantes string
	var cantidadDeIngresantes int

	for i := 0; i < len(contadorDeCarreras); i++ {
		if contadorDeCarreras[i] > cantidadDeIngresantes {
			cantidadDeIngresantes = contadorDeCarreras[i]
			switch i {
			case 0:
				carreraConMasIngresantes = "APU"
			case 1:
				carreraConMasIngresantes = "LI"
			case 2:
				carreraConMasIngresantes = "LS"
			}
		}
	}
	fmt.Println("La carrera con mas ingresantes es: ", carreraConMasIngresantes)
}

func imprimirAnioDeNacimientoConMasIngresantes(aniosDeNacimiento []int) {
	var anioDeNacimientoConMasIngresantes int
	var cantidadDeIngresantes int

	for i := 0; i < len(aniosDeNacimiento); i++ {
		if aniosDeNacimiento[i] > cantidadDeIngresantes {
			cantidadDeIngresantes = aniosDeNacimiento[i]
			anioDeNacimientoConMasIngresantes = aniosDeNacimiento[i]
		}
	}
	fmt.Println("El año de nacimiento con mas ingresantes es: ", anioDeNacimientoConMasIngresantes)
}