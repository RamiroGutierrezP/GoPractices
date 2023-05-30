package main

import "fmt"

func main() {

	contador := 0
	total := 0
	num := 0

	for contador < 250 {
		if num%2 == 0 {
			total += num
			contador++
		}
		num++
	}

	fmt.Println("La suma de los primeros 250 números positivos pares es:", total)
}
