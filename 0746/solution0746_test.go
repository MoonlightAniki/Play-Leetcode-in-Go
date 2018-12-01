package _746

import (
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost []int
		res  int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		if actual := minCostClimbingStairs(tt.cost); actual != tt.res {
			t.Errorf("minCostClimbingStairs(%v), got %d, expected %d\n", tt.cost, actual, tt.res)
		}
	}
}
