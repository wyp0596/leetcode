package leetcode

import (
	"testing"
)

// Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
// n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
// Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
// Note: You may not slant the container and n is at least 2.
// https://leetcode.com/problems/container-with-most-water/#/description
func maxArea(height []int) int {

	// 最大面积
	max := 0
	// 左坐标
	l := 0
	// 右坐标
	r := len(height) - 1

	for l < r {
		// 默认右边比较低
		value := (r - l) * height[r]
		// 右边比较低或者相等
		if height[r] <= height[l] {
			r--
		} else {
			// 左边比较低
			value = (r - l) * height[l]
			l++
		}
		// 判断是否最大面积
		if max < value {
			max = value
		}
	}

	return max
}

func Test11(t *testing.T) {

	var lists = []struct {
		height []int
		result int
	}{
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 3},
			1,
		},
		{
			[]int{3, 1, 5},
			6,
		},
		{
			[]int{0, 0},
			0,
		},
		{
			[]int{1, 3, 5},
			3,
		},
	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("maxArea", l.height, "hope output:", l.result)
		result := maxArea(l.height)
		t.Log("maxArea result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
