package leetcode

import (
	"testing"
)

//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
//Example:
//
//Input: "babad"
//
//Output: "bab"
//
//Note: "aba" is also a valid answer.
//Example:
//
//Input: "cbbd"
//
//Output: "bb"
//https://leetcode.com/problems/longest-palindromic-substring/#/description
func longestPalindrome(s string) string {
	var longest string
	var j int
	for i := 0; i < len(s) && len(longest) < (len(s) - i) * 2; i++ {
		// 奇数个回文字符
		for j = 0; i - j >= 0 && i + j < len(s) && s[i - j] == s[i + j]; j++ {}
		if len(s[i - j + 1 : i + j]) > len(longest) {
			longest = s[i - j + 1 : i + j]
		}
		// 偶数个回文字符
		for j = 0; i - j >= 0 && i + j + 1 < len(s) && s[i - j] == s[i + j + 1]; j++ {}
		if len(s[i - j + 1 : i + j + 1]) > len(longest) {
			longest = s[i - j + 1 : i + j + 1]
		}
	}
	return longest
}

func Test5(t *testing.T) {

	var lists = []struct {
		s      string
		result string
	}{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"",
			"",
		},
		{
			"abababa",
			"abababa",
		},
		{
			"abacbabcabd",
			"bacbabcab",
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("longestPalindrome", l.s, "hope output:", l.result)
		result := longestPalindrome(l.s)
		t.Log("longestPalindrome result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
