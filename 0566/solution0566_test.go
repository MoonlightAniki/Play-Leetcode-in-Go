package solution0566

import (
	"testing"
	"reflect"
)

func Test_matrixReshape(t *testing.T) {
	tests := []struct {
		nums [][]int
		r    int
		c    int
		res  [][]int
	}{
		{
			[][]int{{1, 2}, {3, 4}},
			1,
			4,
			[][]int{{1, 2, 3, 4}},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			2,
			4,
			[][]int{{1, 2}, {3, 4}},
		},
	}

	for _, tt := range tests {
		if actual := matrixReshape(tt.nums, tt.r, tt.c); !reflect.DeepEqual(actual, tt.res) {
			t.Errorf("matrixReshape, got %v, expected %v", actual, tt.res)
		}
	}
}
