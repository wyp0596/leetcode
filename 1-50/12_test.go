package leetcode

import (
	"testing"
)

//Given an integer, convert it to a roman numeral.
//
//Input is guaranteed to be within the range from 1 to 3999.
// https://leetcode.com/problems/integer-to-roman/description/
func intToRoman(num int) string {
	result := ""
	if num > 3999 {
		return result
	}
	if num > 1000 {
		tmp := num / 1000
		num = num % 1000
		for i := 0; i < tmp; i++ {
			result += "M"
		}
	}

	if num > 100 {
		tmp := num / 100
		num = num % 100
		result += intToRomanHelper(tmp, "C", "D", "M")
	}
	if num > 10 {
		tmp := num / 10
		num = num % 10
		result += intToRomanHelper(tmp, "X", "L", "C")
	}
	if num > 0 {
		tmp := num
		result += intToRomanHelper(tmp, "I", "V", "X")
	}
	return result
}

func intToRomanHelper(tmp int, one string, five string, ten string) string {
	result := ""
	if tmp < 4 { // tmp = 1~3
		for i := 0; i < tmp; i++ {
			result += one
		}
	} else if tmp == 4 { // tmp = 4
		result += one
		result += five
	} else if tmp < 9 { // tmp = 5~8
		result += five
		tmp = tmp - 5
		for i := 0; i < tmp; i++ {
			result += one
		}
	} else if tmp == 9 { // tmp = 9
		result += one
		result += ten
	} else { // tmp = 10
		result += ten
	}
	return result
}

func Test12(t *testing.T) {

	var lists = []struct {
		num    int
		result string
	}{
		{
			1,
			"I",
		},
		{
			2,
			"II",
		},
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			5,
			"V",
		},
		{
			6,
			"VI",
		},
		{
			8,
			"VIII",
		},
		{
			9,
			"IX",
		},
		{
			10,
			"X",
		},
		{
			100,
			"C",
		},
		{
			1000,
			"M",
		},
		{
			29,
			"XXIX",
		},
		{
			3999,
			"MMMCMXCIX",
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("intToRoman", l.num, "hope output:", l.result)
		result := intToRoman(l.num)
		t.Log("intToRoman result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
