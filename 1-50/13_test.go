package leetcode

import (
	"testing"
)

//Given a roman numeral, convert it to an integer.
//
//Input is guaranteed to be within the range from 1 to 3999.
//https://leetcode.com/problems/roman-to-integer/#/description
func romanToInt(s string) int {

	//Symbol	I	V	X	L	C	D	M
	//Value		1	5	10	50	100	500	1,000

	//Number	4	9	40	90	400	900
	//Notation	IV	IX	XL	XC	CD	CM

	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		v := s[i]
		switch v {
		case 'M':
			r += 1000
			break
		case 'D':
			r += 500
			break
		case 'C':
			if r >= 500 {
				r -= 100
			} else {
				r += 100
			}
			break
		case 'L':
			r += 50
			break
		case 'X':
			if r >= 50 {
				r -= 10
			} else {
				r += 10
			}
			break
		case 'V':
			r += 5
			break
		case 'I':
			if r >= 5 {
				r -= 1
			} else {
				r += 1
			}
			break
		default:
			panic(string(s[i]))
		}
	}

	return r
}

func Test13(t *testing.T) {

	var lists = []struct {
		s      string
		result int
	}{
		{
			"I",
			1,
		},
		{
			"II",
			2,
		},
		{
			"III",
			3,
		},
		{
			"IV",
			4,
		},
		{
			"V",
			5,
		},
		{
			"VIII",
			8,
		},
		{
			"IX",
			9,
		},
		{
			"MMMCCCXXXIX",
			3339,
		},
		{
			"MMMIM",
			3999,
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("romanToInt", l.s, "hope output:", l.result)
		result := romanToInt(l.s)
		t.Log("romanToInt result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
