package leetcode

import (
	"testing"
	"strings"
)

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
//https://leetcode.com/problems/valid-parentheses/#/description
func isValid(s string) bool {
	chars := strings.Split(s, "")
	var stack []string
	for _, c := range chars {
		if len(stack) == 0 && (c == "}" || c == "]" || c == ")") {
			return false
		}
		if c == "{" || c == "[" || c == "(" {
			stack = append(stack, c)
			continue
		}
		switch stack[len(stack) - 1] {
		case "{":
			if c != "}" {
				return false
			}
			break
		case "[":
			if c != "]" {
				return false
			}
			break
		case "(":
			if c != ")" {
				return false
			}
			break
		default:
			return false
		}
		stack = stack[:len(stack) - 1]
	}

	return len(stack) == 0
}

func Test20(t *testing.T) {

	var lists = []struct {
		s      string
		result bool
	}{
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"(]",
			false,
		},
		{
			"([)]",
			false,
		},
		{
			"([{}])",
			true,
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("isValid", l.s, "hope output:", l.result)
		result := isValid(l.s)
		t.Log("isValid result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
