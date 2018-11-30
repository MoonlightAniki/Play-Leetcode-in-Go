package solution0637_v2

import (
	"testing"
	"reflect"
)

func Test_averageOfLevels(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	tests := []struct {
		root *TreeNode
		res  []float64
	}{
		{root, []float64{3, 14.5, 11}},
	}

	for _, tt := range tests {
		if actual := averageOfLevels(tt.root); !reflect.DeepEqual(actual, tt.res) {
			t.Errorf("averageOfLevels(), got %v, expected %v\n", actual, tt.res)
		}
	}
}

func Benchmark_averageOfLevels(b *testing.B) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	res := []float64{3, 14.5, 11}

	for i := 0; i < b.N; i++ {
		if actual := averageOfLevels(root); !reflect.DeepEqual(actual, res) {
			b.Errorf("averageOfLevels(), got %v, expected %v\n", actual, res)
		}
	}
}
