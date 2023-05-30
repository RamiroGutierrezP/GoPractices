package main

import "fmt"

type PuntoCardinal int

const (
	Norte PuntoCardinal = iota
	Este
	Sur
	Oeste
	Noreste
	Sureste
	Suroeste
	Noroeste
)

func main() {
	var direccion PuntoCardinal
	fmt.Println("Ingrese la dirección del viento: N, S, E, O, NE, SE, SO, NO")
	fmt.Scan(&direccion)
	determinarDireccionDelViente(direccion)
}

func determinarDireccionDelViente (direccion PuntoCardinal) {
	switch direccion {
	case Norte:
		fmt.Println("Norte")
	case Este:
		fmt.Println("Este")
	case Sur:
		fmt.Println("Sur")
	case Oeste:
		fmt.Println("Oeste")
	case Noreste:
		fmt.Println("Noreste")
	case Sureste:
		fmt.Println("Sureste")
	case Suroeste:
		fmt.Println("Suroeste")
	case Noroeste:
		fmt.Println("Noroeste")
	default: 
		fmt.Println("No es una dirección válida")
	}
}
