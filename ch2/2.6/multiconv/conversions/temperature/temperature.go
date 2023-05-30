package temperature

type Celsius float32
type Fahrenheit float32
type Kelvin float32

func (c Celsius) CToF() Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func (c Celsius) CToK() Kelvin     { return Kelvin(c + 273.15) }

func (f Fahrenheit) FToC() Celsius { return Celsius((f - 32) * 5 / 9) }
func (f Fahrenheit) FToK() Kelvin  { return Kelvin((f-32)*5/9 + 273.15) }

func (k Kelvin) KToC() Celsius    { return Celsius(k - 273.15) }
func (k Kelvin) KToF() Fahrenheit { return Fahrenheit((k-273.15)*5/9 + 32) }
