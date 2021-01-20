package _0628

import "sort"

func maximumProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(nums[n-1]*nums[n-2]*nums[n-3], nums[0]*nums[1]*nums[n-1])
}
