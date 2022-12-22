package calculator

import (
	"testing"
)

func TestSuccessfulAdditionOfTwoPositiveNumbers(t *testing.T) {
	result := add(10, 2)
	if result != 12 {
		t.Error("Addition Failed")
	}
}

func TestSuccessfulAdditionOfTwoNegativeNumbers(t *testing.T) {
	result := add(-12, -10)
	if result != -22 {
		t.Error("Addition Failed")
	}
}

func TestSuccessfulAdditionOfTwoIntegers(t *testing.T) {
	result := add(-10, 10)
	if result != 0 {
		t.Error("Addition Failed")
	}
}

func TestSuccessfulSubtractionOfTwoNumbers(t *testing.T) {
	result := subtrat(10, 5)
	if result != 5 {
		t.Error("Subtraction Failed")
	}
}

func TestSuccessfulSubtractionOfTwoNegativeNumbers(t *testing.T) {
	result := subtrat(-10, -5)
	if result != -5 {
		t.Error("Subtraction Failed")
	}
}

func TestSuccessfulSubtractionOfTwoIntegers(t *testing.T) {
	result := subtrat(-10, 5)
	if result != -15 {
		t.Error("Subtraction Failed")
	}
}

func TestSuccessfulMultiplicationOfTwoNumbers(t *testing.T) {
	result := Multiply(10, 30)
	if result != 300 {
		t.Error("Multiplication Failed")
	}
}

func TestSuccessfulMultiplicationOfTwoNegativeNumbers(t *testing.T) {
	result := Multiply(-10, -30)
	if result != 300 {
		t.Error("Multiplication Failed")
	}
}

func TestSuccessfulMultiplicationOfTwoIntegers(t *testing.T) {
	result := Multiply(-10, 30)
	if result != -300 {
		t.Error("Multiplication Failed")
	}
}
