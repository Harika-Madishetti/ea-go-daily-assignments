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
