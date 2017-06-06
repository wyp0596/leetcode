package leetcode

import (
	"testing"
)

//Implement regular expression matching with support for '.' and '*'.
//
//'.' Matches any single character.
//'*' Matches zero or more of the preceding element.
//
//The matching should cover the entire input string (not partial).
//
//The function prototype should be:
//bool isMatch(const char *s, const char *p)
//
//Some examples:
//isMatch("aa","a") → false
//isMatch("aa","aa") → true
//isMatch("aaa","aa") → false
//isMatch("aa", "a*") → true
//isMatch("aa", ".*") → true
//isMatch("ab", ".*") → true
//isMatch("aab", "c*a*b") → true
//https://leetcode.com/problems/regular-expression-matching/#/description
func isMatch(s string, p string) bool {

	// 动态规划
	// 初始化状态表
	dp := make([][]bool, len(s) + 1)
	for i := range dp {
		dp[i] = make([]bool, len(p) + 1)
	}
	// 初始状态,s和p都为空字符串的情况
	dp[0][0] = true

	// 当s为空字符串
	for i := 2; i < len(dp[0]); i++ {
		if p[i - 1] == '*' {
			dp[0][i] = dp[0][i - 2]
		}
	}

	// 从s的第一项开始匹配
	for i := 1; i <= len(s); i++ {
		// 从p的第一项开始匹配
		for j := 1; j <= len(p); j++ {
			if p[j - 1] == '*' {
				dp[i][j] = dp[i][j - 2] || dp[i][j - 1] ||
					(dp[i - 1][j] && (p[j - 2] == '.' || p[j - 2] == s[i - 1]))
			} else {
				dp[i][j] = dp[i - 1][j - 1] && (p[j - 1] == '.' || p[j - 1] == s[i - 1])
			}
		}
	}
	return dp[len(s)][len(p)]
}

func Test10(t *testing.T) {

	var lists = []struct {
		s      string
		p      string
		result bool
	}{
		{
			"a",
			".*..a*",
			false,
		},
		{
			"aa",
			"a",
			false,
		},
		{
			"aa",
			"aa",
			true,
		},
		{
			"a",
			"ab*",
			true,
		},
		{
			"aaa",
			"aa",
			false,
		},
		{
			"aaa",
			"aaaa",
			false,
		},
		{
			"aa",
			"a*",
			true,
		},
		{
			"aab",
			"c*a*b*",
			true,
		},
		{
			"aa",
			".*",
			true,
		},
		{
			"ab",
			".*",
			true,
		},
		{
			"aabcb",
			".*c",
			false,
		},
		{
			"aabc",
			".*c",
			true,
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("isMatch", l.s, "and", l.p, "hope output:", l.result)
		result := isMatch(l.s, l.p)
		t.Log("isMatch result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
