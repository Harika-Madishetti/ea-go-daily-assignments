package calculator

import (
	"testing"
)

func TestSuccessfulAdditionOfTwoPositiveNumbers(t *testing.T) {
	result := add(10, 2)
	if result != 12 {
		t.Error("wrong addition")
	}
}
