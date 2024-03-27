package leetCode

import "slices"

// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
//返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

// 用一个检查数组过滤
func removeDuplicates(nums []int) int {
	var checkArray []int
	for _, i := range nums {
		if slices.Contains(checkArray, i) {
			continue
		}
		checkArray = append(checkArray, i)
	}
	// nums也要赋值为checkArray
	nums = nums[:len(checkArray)]
	copy(nums, checkArray)
	return len(checkArray)
}
