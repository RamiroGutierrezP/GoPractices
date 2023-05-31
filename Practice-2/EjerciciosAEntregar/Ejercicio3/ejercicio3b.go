package main

import "fmt"



func main() {
	v:= -23
	b:= 2
	fmt.Println(Convert(v, b))
}


func Convert (v int, b int) string {

	var resultado string
	negativo := false

	//Verifico que la base sea mayor que 1 y menor que 37
	if b < 1 || b > 37 {
		return "Base invalida"
	}
	//Si v es negativo, lo paso a positivo para poder operar
	if v < 0 {
		v = v * -1
		negativo = true
	}

	//Este string lo voy a usar como índice para obtener el valor en la base
	digitos := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//Itero hasta que el valor sea 0 o menor
	for v > 0 {
		resultado = string(digitos[v%b]) + resultado
		v = v / b
	}
	//Vuelvo a chequear si es negativo para agregarle el signo
	if negativo {
		resultado = "-" + resultado
	}
	return resultado
}