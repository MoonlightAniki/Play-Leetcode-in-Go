package _387

import "testing"

func Test_firstUniqChar(t *testing.T) {
	tests := []struct {
		s   string
		res int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
	}

	for _, tt := range tests {
		if actual := firstUniqChar(tt.s); actual != tt.res {
			t.Errorf("firstUniqChar(%s), got %d, expected %d\n", tt.s, actual, tt.res)
		}
	}
}
