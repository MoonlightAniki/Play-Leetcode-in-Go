package solution0485

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums []int
		res  int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 1, 1, 1, 1, 1}, 6},
		{[]int{0, 0, 0, 0, 0, 0}, 0},
		{[]int{1, 1, 1, 0, 0, 1}, 3},
	}

	for _, tt := range tests {
		if actual := findMaxConsecutiveOnes(tt.nums); actual != tt.res {
			t.Errorf("findMaxConsecutiveOnes(%v), got %d, expected %d\n", tt.nums, actual, tt.res)
		}
	}
}
