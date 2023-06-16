package main

import "fmt"

type Map[K comparable, V any] map[K]V

func main() {

	m := make(Map[string, int])

	x := make(Map[int, string])

	m["ramiro"] = 5
	m["pulga"] = 10	
	m["chapa"] = 8

	x[5] = "ramiro"
	x[10] = "pulga"	
	x[8] = "chapa"

	fmt.Println(m, x)
}