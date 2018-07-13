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
		result := Add(tt.left, tt.right)
		if result != tt.expected {
			t.Errorf("%d is expected. got=%d", tt.expected, result)
		}
	}
}
