package main

import (
	"fmt"
	"math"
)

func main() {
	var num int

	fmt.Println("Ingrese un número: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Error al leer el número:", err)
		return
	}
	fmt.Println(funcion(num))
}

func funcion(num int) int {
	switch {
	case num < -18:
		return int(math.Abs(float64(num)))
	case num >= -18 && num <= -1:
		return num % 4
	case num >= 1 && num < 20:
		return int(math.Pow(float64(num), 2))
	case num >= 20:
		return num * (-1)
	}
	return 0
}

// Lo unico que hago es sacar los casteos a int
func funcionFlotante(num float64) float64 {
	switch {
	case num < -18:
		return math.Abs(num)
	case num >= -18 && num <= -1:
		return math.Mod(num, 4)
	case num >= 1 && num < 20:
		return math.Pow(num, 2)
	case num >= 20:
		return num * (-1)
	}
	return 0
}
