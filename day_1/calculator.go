package calculator

import "errors"

var errorDivisionByZero = errors.New("Divided By Zero")

func add(x float64, y float64) float64 {
	return x + y
}

func subtrat(x float64, y float64) float64 {
	return x - y
}

func Multiply(x float64, y float64) float64 {
	return x * y
}

func Divide(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, errorDivisionByZero
	}
	return x / y, nil
}
