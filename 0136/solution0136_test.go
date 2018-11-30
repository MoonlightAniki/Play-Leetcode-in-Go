package solution0136

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		res  int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tests {
		if actual := singleNumber(tt.nums); actual != tt.res {
			t.Errorf("singleNumber(%v), got %d, expected %d\n", tt.nums, actual, tt.res)
		}
	}
}

func Benchmark_singleNumber(b *testing.B) {
	nums, res := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 10, 10}, 9
	for i := 0; i < b.N; i++ {
		if actual := singleNumber(nums); actual != res {
			b.Errorf("singleNumber(%v), got %d, expected %d\n", nums, actual, res)
		}
	}
}
