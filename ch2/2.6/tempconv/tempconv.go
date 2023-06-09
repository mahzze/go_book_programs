package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (kelvin Kelvin) String() string     { return fmt.Sprintf("%gK", kelvin) }
func (celsius Celsius) String() string   { return fmt.Sprintf("%g°C", celsius) }
func (fahren Fahrenheit) String() string { return fmt.Sprintf("%g°F", fahren) }
