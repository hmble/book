package tempconv

func CtoF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }

func FtoC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
