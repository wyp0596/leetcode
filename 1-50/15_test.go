package leetcode

import (
	"testing"
	"sort"
	"reflect"
)

//Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
//
//Note: The solution set must not contain duplicate triplets.
//
//For example, given array S = [-1, 0, 1, 2, -1, -4],
//
//A solution set is:
//[
//	[-1, 0, 1],
//	[-1, -1, 2]
//]
//https://leetcode.com/problems/3sum/#/description
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	r := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		lo := i + 1
		hi := len(nums) - 1
		sum := 0 - nums[i]

		for lo < hi {
			if nums[lo] + nums[hi] == sum {
				r = append(r, []int{nums[i], nums[lo], nums[hi]})
				for lo < hi && nums[lo] == nums[lo + 1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi - 1] {
					hi--
				}
				lo++
				hi--
			} else if nums[lo] + nums[hi] < sum {
				lo++
			} else {
				hi--
			}
		}
		for i < len(nums) - 2 && nums[i] == nums[i+1] {
			i++
		}
	}

	return r
}

func Test15(t *testing.T) {
	var lists = []struct {
		nums   []int
		result [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
		},
		{
			[]int{0, 0, 0},
			[][]int{[]int{0, 0, 0}},
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("threeSum", l.nums, "hope output:", l.result)
		result := threeSum(l.nums)
		if !reflect.DeepEqual(l.result, result) {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
