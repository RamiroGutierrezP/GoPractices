package main

import "fmt"

func main() {
	fmt.Println("Factorial of 5 is", factorial(5))
}


func factorial(n int) int {

	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}