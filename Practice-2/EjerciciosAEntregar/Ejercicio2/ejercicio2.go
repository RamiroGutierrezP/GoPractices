package main

import (
	"bufio"
	"fmt"
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
	var contadorDeCarreras [3]int
	ingresantes := cargarIngresantes()
	barilochenses := make([]string, 0)
	aniosDeNacimiento := make(map[int]int)

	//Recorro el slice de ingresantes y voy procesando los datos
	for i := 0; i < len(ingresantes); i++ {
		
		if ingresantes[i].presentoTitulo == false {
			ingresantes = append(ingresantes[:i], ingresantes[i+1:]...)
			i--
			continue
		} else {
			if (ingresantes[i].ciudadOrigen == "Bariloche") {
				barilochenses = append(barilochenses, ingresantes[i].nombre + " " + ingresantes[i].apellido)
			}
			contadorDeCarreras = incrementarContador(ingresantes[i].codigoDeCarrera, contadorDeCarreras)
			aniosDeNacimiento[ingresantes[i].fechaNacimiento.anio]++
		}
	}
	
	//Informo los resultados 
	informarCarreraConMasIngresantes(contadorDeCarreras)
	imprimirBarilochenses(barilochenses)
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
func informarCarreraConMasIngresantes (contadorDeCarreras [3]int) {
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
func cargarIngresantes () []Ingresante {

	ingresantes := make([]Ingresante, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var ingresante Ingresante

		linea := scanner.Text()
		datos := strings.Split(linea, ",")

		// Valido que la linea tenga 8 datos; si no tiene salteo la linea
		if len(datos) != 8 {
			continue
		}

		ingresante.apellido = datos[0]
		ingresante.nombre = datos[1]
		ingresante.ciudadOrigen = datos[2]
		ingresante.fechaNacimiento.dia = convertirStringAInt(datos[3])
		ingresante.fechaNacimiento.mes = convertirStringAInt(datos[4])
		ingresante.fechaNacimiento.anio = convertirStringAInt(datos[5])
		ingresante.presentoTitulo = convertirStringABool(datos[6])
		ingresante.codigoDeCarrera = datos[7]

		ingresantes = append(ingresantes, ingresante)
	}
	return ingresantes
}
func imprimirBarilochenses (barilochenses []string) {
	fmt.Println("Los ingresantes de Bariloche son: ")
	for i := 0; i < len(barilochenses); i++ {
		fmt.Println(barilochenses[i])
	}
}

func imprimirAnioDeNacimientoConMasIngresantes (aniosDeNacimiento map[int]int) {
	var anioMaximo int
	var maximo int

	for anio, cantidad := range aniosDeNacimiento {
		if cantidad > maximo {
			maximo = cantidad
			anioMaximo = anio
		}
	}
	fmt.Println("El año de nacimiento con mas ingresantes es: ", anioMaximo)
}