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
	var acumulator [4]float64
	var counterHigh, counterMedium, counterLow int

	//Iterate 10 times reading the temperatures
	for i := 0; i < 10; i++ {
		scanner.Scan();
		fmt.Sscan(scanner.Text(), &temperature);

		if temperature < 36 && temperature > 20{
			acumulator[0] += temperature
			counterLow++
		} else if temperature >= 36 && temperature <= 37.5 {
			acumulator[1] += temperature
			counterMedium++
		} else if temperature > 37.5 && temperature < 50 {
			acumulator[2] += temperature
			counterHigh++
		} else {
			acumulator[3] += temperature
		}
	}

	//Print the average temperature of each group
	fmt.Println("Average temperature of high gruop: ", acumulator[2]/float64(counterHigh))
	fmt.Println("Average temperature of medium gruop: ", acumulator[1]/float64(counterMedium))
	fmt.Println("Average temperature of low gruop: ", acumulator[0]/float64(counterLow))
	fmt.Println("Total of incorrects temperatures: ", acumulator[3])
}
