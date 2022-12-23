package calculator

import (
	"errors"
	"math"
)

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

func SinFunction(x float64) float64 {
	return math.Sin(x)
}

func CosFunction(x float64) float64 {
	return math.Cos(x)
}
