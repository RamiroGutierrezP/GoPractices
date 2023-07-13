package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	// Leer la frase
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una cadena: ")
	frase, _ := reader.ReadString('\n')

	fraseMin := strings.ToLower(frase)


	for strings.Contains(fraseMin, "automóvil") {

		//Obtengo la posición del primer "automóvil" que se encuentre
		pos := strings.Index(fraseMin, "automóvil")

		//Lo cambio por el "miércoles"
		frase = corregir(frase, pos)

		fraseMin = strings.ToLower(frase)
	}

	// Imprimo
	fmt.Println("Cadena modificada: ", frase)

}

func corregir(frase string, pos int) string {

	//En este caso tengo que hacer explicito el slice de runas porque el caracter "é" ocupa 2 bytes
	miercoles := []rune(frase[pos : pos+10]) 

	reemplazo := []rune("miércoles")

	for i, char := range miercoles {
		if unicode.IsUpper(char) {
			reemplazo[i] = unicode.ToUpper(reemplazo[i])
		}
	}

	return strings.ReplaceAll(frase, string(miercoles), string(reemplazo))
}
