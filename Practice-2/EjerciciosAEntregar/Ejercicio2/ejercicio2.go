package main

import "fmt"

type Ingresante struct {
	apellido string
	nombre string
	ciudadOrigen string
	fechaNacimiento Fecha
	tituloSecundario bool
	codigoDeCarrera string
}

type Fecha struct {
	dia int
	mes int
	anio int
}

func main() {
	fmt.Println("Hello, World!")
}
