package main

import "fmt"

func main() {

	/* i n t e g e r s */
	var zz int = 10
	x := 10
	var z int = x
	var y int = x + 1

	fmt.Println(zz, z, x, y)

	/* c o n s t a n t s */
	const n = 5001
	const cint = 5001
	fmt.Println(n, cint)

	/* f l o a t */
	var e float32 = 6
	ffloat32 := e

	fmt.Println(e, ffloat32)
}
