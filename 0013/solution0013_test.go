package _013

import (
	"testing"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		s string
		r int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, tt := range tests {
		if actual := romanToInt(tt.s); actual != tt.r {
			t.Errorf("romanToInt(%s), got %d, expected %d\n", tt.s, actual, tt.r)
		}
	}
}
