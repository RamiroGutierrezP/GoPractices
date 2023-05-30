package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Create the scanner variable to read the input
	file, err := os.Open("Practice-2/Ejercicio1/input.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//Create the variables
	var temperature float64
	var acumulatorHigh, acumulatorMedium, acumulatorLow float64
	var counterHigh, counterMedium, counterLow int

	//Iterate 10 times reading the temperatures
	for i := 0; i < 10; i++ {
		scanner.Scan();
		fmt.Sscan(scanner.Text(), &temperature);

		if temperature < 36 {
			acumulatorLow += temperature
			counterLow++
		} else if temperature >= 36 && temperature <= 37.5 {
			acumulatorMedium += temperature
			counterMedium++
		} else {
			acumulatorHigh += temperature
			counterHigh++
		}
	}

	//Print the average temperature of each group
	fmt.Println("Average temperature of high gruop: ", acumulatorHigh/float64(counterHigh))
	fmt.Println("Average temperature of medium gruop: ", acumulatorMedium/float64(counterMedium))
	fmt.Println("Average temperature of low gruop: ", acumulatorLow/float64(counterLow))
}
