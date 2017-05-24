package leetcode

import (
	"testing"
	"math"
	"strings"
)

//Implement atoi to convert a string to an integer.
//
//Hint: Carefully consider all possible input cases.
// If you want a challenge, please do not see below and ask yourself what are the possible input cases.
//
//Notes: It is intended for this problem to be specified vaguely (ie, no given input specs).
// You are responsible to gather all the input requirements up front.
//https://leetcode.com/problems/string-to-integer-atoi/#/description
func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}
	negative := false
	first := str[0]
	if first == '-' {
		negative = true
		str = str[1:]
	} else if first == '+' {
		str = str[1:]
	} else if first < '0' || first > '9' {
		return 0
	}

	result := 0
	for _, r := range str {
		x := int(r - '0')
		if x < 0 || x > 9 {
			break
		}
		result = result * 10 + x
		if result > math.MaxInt32 {
			if negative {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	if negative {
		return -result
	}
	return result
}

func Test8(t *testing.T) {

	var lists = []struct {
		s      string
		result int
	}{
		{
			"",
			0,
		},
		{
			"   123 ",
			123,
		},
		{
			"+123456",
			123456,
		},
		{
			"-123456",
			-123456,
		},
		{
			"+21474836489",
			2147483647,
		},
		{
			"-21474836489",
			-2147483648,
		},
		{
			"214748a36489",
			214748,
		},
		{
			"-214748a36489",
			-214748,
		},
		{
			"+21474836489a123",
			2147483647,
		},
		{
			"-2147483648c1239",
			-2147483648,
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("myAtoi", l.s, "hope output:", l.result)
		result := myAtoi(l.s)
		t.Log("myAtoi result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
