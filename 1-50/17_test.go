package leetcode

import (
	"testing"
	"strings"
	"reflect"
)

//Given a digit string, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below.
//
//
//
//Input:Digit string "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//Note:
//Although the above answer is in lexicographical order, your answer could be in any order you want.
//https://leetcode.com/problems/letter-combinations-of-a-phone-number/#/description
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	maps := map[string][]string{
		"2":[]string{"a", "b", "c"},
		"3":[]string{"d", "e", "f"},
		"4":[]string{"g", "h", "i"},
		"5":[]string{"j", "k", "l"},
		"6":[]string{"m", "n", "o"},
		"7":[]string{"p", "q", "r", "s"},
		"8":[]string{"t", "u", "v"},
		"9":[]string{"w", "x", "y", "z"},
	}
	result := []string{""}
	digitList := strings.Split(digits, "")
	for _, d := range digitList {
		minus := len(result)
		letters := maps[d]
		for _, r := range result {
			for _, l := range letters {
				result = append(result, r + l)
			}
		}
		result = result[minus:]
	}
	return result
}

func Test17(t *testing.T) {

	var lists = []struct {
		s      string
		result []string
	}{
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("letterCombinations", l.s, "hope output:", l.result)
		result := letterCombinations(l.s)
		t.Log("letterCombinations result:", result)
		if !reflect.DeepEqual(l.result, result) {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
