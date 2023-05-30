package main

import "fmt"

func main() {

	var car rune

	fmt.Println("Ingrese una sigla de un punto cardinal: ")
	_, err := fmt.Scanf("%c", &car)
	if err != nil {
		fmt.Println("Error al leer el número:", err)
		return
	}
	fmt.Println(funcion(car))
}

func funcion(car rune) string {
	switch car {
	case 'N':
		return "Usted se dirige hacia el Norte"
	case 'O':
		return "Usted se dirige hacia el Oeste"
	case 'S':
		return "Usted se dirige hacia el Sur"
	case 'E':
		return "Usted se dirige hacia el Este"
	default:
		return "El valor ingresado no corresponde a un punto cardinal"
	}
}
