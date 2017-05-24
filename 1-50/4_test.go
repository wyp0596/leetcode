package leetcode

import (
	"testing"
)

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//Example 1:
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//Example 2:
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5
//https://leetcode.com/problems/median-of-two-sorted-arrays/#/description
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if nums1 == nil && nums2 == nil {
		return -1.0
	}
	median := []int{}
	totalLen := len(nums1) + len(nums2)
	halfLen := totalLen / 2;
	target1 := halfLen
	target2 := halfLen
	if totalLen % 2 == 0 {
		target1 = halfLen - 1
	}
	index := 0
	i := 0
	j := 0
	for {
		var value int
		if i == len(nums1) {
			value = nums2[j]
			j++
		} else if j == len(nums2) {
			value = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			value = nums1[i]
			i++
		} else {
			value = nums2[j]
			j++
		}
		if index == target1 || index == target2 {
			median = append(median, value)
		}
		index++
		if index > target2 {
			break
		}
	}
	sum := 0
	for _, v := range median {
		sum += v
	}
	return float64(sum) / float64(len(median))
}

func Test4(t *testing.T) {

	var lists = []struct {
		nums1  []int
		nums2  []int
		result float64
	}{
		{
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		{
			nil,
			nil,
			-1.0,
		},
		{
			[]int{1, 3, 5, 7},
			nil,
			4.0,
		},
		{
			nil,
			[]int{1, 3, 5, 7},
			4.0,
		},

	}
	t.Log("Begin Tests")
	for i, l := range lists {
		t.Log("begin test #", i)
		t.Log("findMedianSortedArrays", l.nums1, "and", l.nums2, "hope output:", l.result)
		result := findMedianSortedArrays(l.nums1, l.nums2)
		t.Log("findMedianSortedArrays result:", result)
		if l.result != result {
			t.Fatal("result error, want", l.result, "but output", result)
		}
		t.Log("end test #", i)
	}
	t.Log("End Tests")

}
