package leetcode

import (
	"testing"
	"bytes"
)

//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//And then read line by line: "PAHNAPLSIIGYIR"
//Write the code that will take a string and make this conversion given a number of rows:
//
//string convert(string text, int nRows);
//convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
//
//https://leetcode.com/problems/zigzag-conversion/#/description
func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= numRows {
		return s
	}
	res := bytes.Buffer{}
	magicNum := numRows*2 - 2
	specialNum := magicNum
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); {
			res.WriteByte(s[j])
			if specialNum < magicNum && specialNum > 0 && j+specialNum < len(s) {
				res.WriteByte(s[j+specialNum])
			}
			j += magicNum
		}
		specialNum -= 2
	}

	return res.String()
}

func Test6(t *testing.T) {

	var lists = []struct {
		s       string
		numRows int
		result  string
	}{
		{
			"PAYPALISHIRING",
			0,
			"PAYPALISHIRING",
		},
		{
			"PAYPALISHIRING",
			1,
			"PAYPALISHIRING",
		},
		{
			"PAYPALISHIRING",
			2,
			"PYAIHRNAPLSIIG",
		},
		{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("convert", l.s, "numRows", l.numRows, "hope output:", l.result)
		result := convert(l.s, l.numRows)
		t.Log("convert result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
