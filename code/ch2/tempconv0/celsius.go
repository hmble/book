package tempconv

import "fmt"

type Celsius float64
type Farenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }

func FtoC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g C", c) }
