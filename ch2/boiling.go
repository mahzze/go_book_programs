package main

import "fmt"

const boilingF = 212.0

func main() {
	var fahrenheit = boilingF
	var celsius = (fahrenheit - 32) * 5 / 9
	fmt.Printf("boiling point = %gºF or %gºC\n", fahrenheit, celsius)
}
