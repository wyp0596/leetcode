package leetcode

import (
	"testing"
	"reflect"
)
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
//https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	for i1, v1 := range nums {
		for i2, v2 := range nums {
			if i1 < i2 && v1 + v2 == target {
				return []int{i1, i2}
			}
		}
	}
	return nil
}

func Test1(t *testing.T) {

	var lists = []struct {
		nums   []int
		target int
		result []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{-5, 4, 3, 6},
			1,
			[]int{0, 3},
		},
		{
			[]int{0, 4, 3, 6},
			7,
			[]int{1, 2},
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("input:", l.nums, "target:", l.target)
		result := twoSum(l.nums, l.target)
		t.Log("twoSum result:", result)
		if !reflect.DeepEqual(result, l.result) {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
