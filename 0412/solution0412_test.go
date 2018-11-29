package solution0412

import (
	"testing"
	"reflect"
)

// 表格驱动测试
// go test .
// go test -coverprofile=c.out  将代码覆盖率报告输出到c.out中
// go tool cover -html=c.out    在网页中查看代码覆盖率报告
func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n   int
		res []string
	}{
		{1, []string{"1"}},
		{2, []string{"1", "2"}},
		{3, []string{"1", "2", "Fizz"}},
		{4, []string{"1", "2", "Fizz", "4"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, tt := range tests {
		if actual := fizzBuzz(tt.n); !reflect.DeepEqual(actual, tt.res) {
			t.Errorf("fizzBuzz(%d), got %v, expected %v", tt.n, actual, tt.res)
		}
	}
}

// 性能测试
// go test -bench .
// go test -bench -cpuprofile cpu.out   将性能测试报告输出到cpu.out中
// go tool pprof cpu.out                查看性能测试报告
func BenchmarkFizzBuzz(b *testing.B) {
	n, res := 15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	for i := 0; i < b.N; i++ {
		actual := fizzBuzz(n)
		if !reflect.DeepEqual(actual, res) {
			b.Errorf("fizzBuzz(%d), got %v, expected %v", n, actual, res)
		}
	}
}
