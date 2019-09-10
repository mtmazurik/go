package temperature

func FtoC(f float64) float64 {
	return (f - 32.0/9.0) * 5.0
}

func CtoF(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}
