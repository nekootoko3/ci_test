package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{1, 2, 3},
		{-1, 2, 1},
		{0, 0, 0},
	}

	for _, tt := range tests {
		sum := Add(tt.left, tt.right)
		if sum != tt.expected {
			t.Errorf("%d is expected. got=%d", tt.expected, sum)
		}
	}
}

func TestSubstract(t *testing.T) {
	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{10, 1, 9},
		{0, 0, 0},
		{-4, -1, -3},
	}

	for _, tt := range tests {
		diff := Substract(tt.left, tt.right)
		if diff != tt.expected {
			t.Errorf("%d is expected. got=%d", tt.expected, diff)
		}
	}
}
