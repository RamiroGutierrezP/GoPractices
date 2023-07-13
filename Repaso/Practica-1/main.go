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

	//Creo un índice para recorrer la frase
	i := 0
	
	for (i + 6) < len(frase) {
		//Pregunto si la letra es "j" o "J"
		if (frase[i] == 'j' || frase[i] == 'J') {
			//Si es J o j agarro las siguientes 6 letras
			candidato := frase[i : i+6]
			if (strings.ToLower(candidato) == "jueves") {
				frase = corregir(frase, candidato)
				i += 6
				continue
			}
		}
		i++
	}

	// Imprimo
	fmt.Println("Frase corregida: ", frase)

}

func corregir(frase string, jueves string) string {

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
