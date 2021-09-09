package tempconv

type Celsius float64
type Fahrenheit float64

//const (
//	AbsoluteZeroC celsius = -273.15
//	FreezingC
//	celsius = 0
//	BoilingC
//	celsius = 100
//)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
