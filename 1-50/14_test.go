package leetcode

import (
	"testing"
	"strings"
)

//Write a function to find the longest common prefix string amongst an array of strings.
//https://leetcode.com/problems/longest-common-prefix/#/description
func longestCommonPrefix(strs []string) string {
	r := ""
	i := 1

	for len(strs) > 0 && len(strs[0]) >= i {
		temp := strs[0][:i]
		for _, s := range strs {
			if strings.Index(s, temp) != 0 {
				return r
			}
		}
		r = temp
		i++
	}
	return r
}

func Test14(t *testing.T) {
	var lists = []struct {
		strs   []string
		result string
	}{
		{
			[]string{},
			"",
		},
		{
			[]string{"ab"},
			"ab",
		},
		{
			[]string{"ab", "abc"},
			"ab",
		},
		{
			[]string{"abcdefg", "a"},
			"a",
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("longestCommonPrefix", l.strs, "hope output:", l.result)
		result := longestCommonPrefix(l.strs)
		t.Log("longestCommonPrefix result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
