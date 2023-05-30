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

	//Iterate 10 times reading the temperatures
	for i := 0; i < 10; i++ {
		scanner.Scan();
		fmt.Sscan(scanner.Text(), &temperature);

		celsiusToFahrenheit(temperature);
		fmt.Println("Temperature in Fahrenheit: ", celsiusToFahrenheit(temperature));
	}
}


//A function that receive temperatures in Celsius and return the temperature in Fahrenheit

func celsiusToFahrenheit(temperature float64) float64 {
	return temperature * 1.8 + 32
}