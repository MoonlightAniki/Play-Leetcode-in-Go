package solution0448

import (
	"testing"
	"reflect"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		res  []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}
	for _, tt := range tests {
		if actual := findDisappearedNumbers(tt.nums); !reflect.DeepEqual(actual, tt.res) {
			t.Errorf("findDisappearedNumbers(%v), got %v, expected %v\n", tt.nums, actual, tt.res)
		}
	}
}

func Benchmark_findDisappearedNumbers(b *testing.B) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := []int{5, 6}
	for i := 0; i < b.N; i++ {
		if actual := findDisappearedNumbers(nums); !reflect.DeepEqual(actual, res) {
			b.Errorf("findDisappearedNumbers(%v), got %v, expected %v\n", nums, actual, res)
		}
	}
}
