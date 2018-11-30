package solution0917

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	tests := []struct{ s, res string }{
		{"ab-cd", "dc-ba"},
		{"a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
		{"Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},
	}
	for _, tt := range tests {
		if actual := reverseOnlyLetters(tt.s); actual != tt.res {
			t.Errorf("reverseOnlyLetters(%s), got %s, expected %s\n", tt.s, actual, tt.res)
		}
	}
}
