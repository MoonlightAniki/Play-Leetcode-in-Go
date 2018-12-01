package _720

import "testing"

func Test_longestWord(t *testing.T) {
	tests := []struct {
		words []string
		res   string
	}{
		{[]string{"w", "wo", "wor", "worl", "world"}, "world"},
		{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"},
		{[]string{"rac","rs","ra","on","r","otif","o","onpdu","rsf","rs","ot","oti","racy","onpd"}, "otif"},
		{[]string{"b", "br", "bre", "brea", "break", "breakf", "breakfa", "breakfas", "breakfast", "l", "lu", "lun", "lunc", "lunch", "d", "di", "din", "dinn", "dinne", "dinner"}, "breakfast"},
	}
	for _, tt := range tests {
		if actual := longestWord(tt.words); actual != tt.res {
			t.Errorf("longestWord(%v), got %s, expected %s\n", tt.words, actual, tt.res)
		}
	}
}
