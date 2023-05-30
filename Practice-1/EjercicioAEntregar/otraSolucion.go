package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main1() {

	// Leer la frase
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una cadena: ")
	frase, _ := reader.ReadString('\n')

	// Pasar la frase a minúscula y crear una frase cambiando martes por jueves
	fraseMin := strings.ToLower(frase)
	fraseMartes := strings.ReplaceAll(fraseMin, "jueves", "martes")

	// Creo una frase nueva que va a ser mutable para modificar las mayusculas
	fraseFinal := []byte(fraseMartes)

	// Recorro la frase
	for i := 0; i < len(frase); i++ {
		// Si es mayuscula... modifico
		if unicode.IsUpper(rune(frase[i])) {
			fraseFinal[i] = byte(unicode.ToUpper(rune(fraseMartes[i])))
		}
	}

	// Imprimo
	fmt.Println(frase, string(fraseFinal))

}
