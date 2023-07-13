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

	// Pasar la frase a minúscula para localizar mas facil un "jueves"
	fraseMin := strings.ToLower(frase)

	//Mientras la frase tenga "jueves"
	for strings.Contains(fraseMin, "jueves") {

		//Obtengo la posición del primer jueves que se encuentre
		pos := strings.Index(fraseMin, "jueves")

		//Lo cambio por el martes
		frase = corregir(frase, pos)

		//Tengo que cambiarlo tambien acá para que no me de siempre la misma posicion
		fraseMin = strings.ToLower(frase)

	}

	// Imprimo
	fmt.Println("Frase corregida: ", frase)

}

func corregir(frase string, pos int) string {

	// Obtengo el formato del jueves
	jueves := frase[pos : pos+6]

	// Creo una cadena mutable para despues reemplazar el martes
	reemplazo := []rune("martes")

	// Recorro la palabra
	for i, char := range jueves {
		// Si es mayuscula... modifico
		if unicode.IsUpper(char) {
			reemplazo[i] = unicode.ToUpper(reemplazo[i])
		}
	}

	//Remplazo
	return strings.ReplaceAll(frase, jueves, string(reemplazo))
}
