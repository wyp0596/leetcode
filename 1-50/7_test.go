package leetcode

import (
	"testing"
)

//Reverse digits of an integer.
//
//Example1: x = 123, return 321
//Example2: x = -123, return -321
//
//click to show spoilers.
//
//Note:
//The input is assumed to be a 32-bit signed integer.
// Your function should return 0 when the reversed integer overflows.
//https://leetcode.com/problems/reverse-integer/#/description
func reverse(x int) int {
	result := 0
	for x != 0 {
		mod := x % 10
		temp := result * 10 + mod
		result = int(int32(temp))
		if result != temp {
			return 0
		}
		x /= 10
	}
	return result
}

func Test7(t *testing.T) {

	var lists = []struct {
		x      int
		result int
	}{
		{
			123,
			321,
		},
		{
			-123,
			-321,
		},
		{
			2147483647,
			0,
		},
		{
			-2147483647,
			0,
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("reverse", l.x, "hope output:", l.result)
		result := reverse(l.x)
		t.Log("reverse result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
