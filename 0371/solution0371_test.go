package _371

import "testing"

func Test_getSum(t *testing.T) {
	tests := []struct{ a, b, res int }{
		{1, 2, 3},
		{-2, 3, 1},
	}
	for _, tt := range tests {
		if actual := getSum(tt.a, tt.b); actual != tt.res {
			t.Errorf("getSum(%d, %d), got %d, expected %d\n", tt.a, tt.b, actual, tt.res)
		}
	}
}
