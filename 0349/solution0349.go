package _349

// 349. Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/
/*
Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:
Each element in the result must be unique.
The result can be in any order.
 */
func intersection(nums1 []int, nums2 []int) []int {
	m1 := getInts(nums1)
	m2 := getInts(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	res := make([]int, 0)
	// m1 是较短的字典，会快一些
	for k, _ := range m1 {
		if m2[k] {
			res = append(res, k)
		}
	}
	return res
}

func getInts(nums []int) map[int]bool {
	res := make(map[int]bool, len(nums))
	for _, num := range nums {
		res[num] = true
	}
	return res
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Intersection of Two Arrays.
