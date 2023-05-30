package main

import "fmt"

func main() {
	fmt.Println("Factorial of 5 is", recursiveFactorial(5))
}


func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}