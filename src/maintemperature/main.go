// main - for temperature calls the temperature package
package main

import (
	"fmt"
	"strconv"
	"temperature"
)

func main() {
	fmt.Printf("Celsius> ")

	var inputValue string
	fmt.Scanln(&inputValue)

	var celsius float64

	var err error
	celsius, err = strconv.ParseFloat(inputValue, 64)
	if err != nil {
		panic(err)
	}
	fahrenheit := temperature.CtoF(celsius)
	fmt.Printf("Celsius: %f Fahrenheit: %f \n", celsius, fahrenheit)
}
