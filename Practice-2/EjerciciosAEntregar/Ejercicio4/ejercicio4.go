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

	formato := frase[pos : pos+10] 
	palabra := []rune(formato)  
	reemplazo := []rune("miércoles")

	for i := 0; i < 9; i++ {
		if unicode.IsUpper(palabra[i]) {
			reemplazo[i] = unicode.ToUpper(rune(reemplazo[i]))
		}
	}

	return strings.ReplaceAll(frase, string(palabra), string(reemplazo))
}
