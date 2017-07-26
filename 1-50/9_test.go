package leetcode

import (
	"testing"
	"math"
)

//Determine whether an integer is a palindrome. Do this without extra space.
//
//Some hints:
//Could negative integers be palindromes? (ie, -1)
//
//If you are thinking of converting the integer to string, note the restriction of using extra space.
//
//You could also try reversing an integer. However, if you have solved the problem "Reverse Integer",
// you know that the reversed integer might overflow. How would you handle such case?
//
//There is a more generic way of solving this problem.
//https://leetcode.com/problems/palindrome-number/#/description
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	input := x
	y := 0
	for x > 0 {
		y = y*10 + x%10
		if y > math.MaxInt32 {
			return false
		}
		x = x / 10
	}
	return input == y
}

func Test9(t *testing.T) {

	var lists = []struct {
		x      int
		result bool
	}{
		{
			-121,
			false,
		},
		{
			-123,
			false,
		},
		{
			-1,
			false,
		},
		{
			0,
			true,
		},
		{
			1,
			true,
		},
		{
			121,
			true,
		},
		{
			123,
			false,
		},
		{
			2147483646,
			false,
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("isPalindrome", l.x, "hope output:", l.result)
		result := isPalindrome(l.x)
		t.Log("isPalindrome result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
