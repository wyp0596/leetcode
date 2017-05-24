package leetcode

import (
	"testing"
	"strings"
	"unicode/utf8"
)

//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3.
// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
//https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description
func lengthOfLongestSubstring(s string) int {
	var subStr string
	maxLen := 0
	for _, r := range s {
		i := strings.Index(subStr, string(r))
		if i != -1 {
			subStr = subStr[i + 1:]
		}
		subStr += string(r)
		if utf8.RuneCountInString(subStr) > maxLen {
			maxLen = utf8.RuneCountInString(subStr)
		}
	}
	return maxLen
}

func Test3(t *testing.T) {

	var lists = []struct {
		s      string
		result int
	}{
		{
			"abcaab",
			3,
		},
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"ababcde",
			5,
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("lengthOfLongestSubstring", l.s, "hope output:", l.result)
		result := lengthOfLongestSubstring(l.s)
		t.Log("lengthOfLongestSubstring result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
