package tempconv

func CToF(c tempconv.Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f tempconv.Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
