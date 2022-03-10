package algorithm

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var x = 1312
	s := fmt.Sprintf("%d", x)
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-i-1] {
			t.Log(false)
			return
		}
	}
	t.Log(true)
}
