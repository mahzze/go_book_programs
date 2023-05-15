package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func CToF(celsius Celsius) Fahrenheit { return Fahrenheit(celsius*9/5 + 32) }

func FToC(fahren Fahrenheit) Celsius { return Celsius((fahren - 32) * 5 / 9) }

// this is a method, it associates to the type celsius the String() function
func (celsius Celsius) String() string { return fmt.Sprintf("%gºC", celsius) }

func main() {
	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0) //true
	//  fmt.Println(c == f)          // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true", but just because both variables are declared to 0 right now

	c = FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}
