/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	/* Brute Fource */
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[j] == target-nums[i] {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	/* Tow-pass Hash Table */
	// m := make(map[int]int)
	// for i, num := range nums {
	// 	m[num] = i
	// }
	// for i, num := range nums {
	// 	complement := target - num
	// 	if j, ok := m[complement]; ok && j != i {
	// 		return []int{i, j}
	// 	}
	// }

	/* One-pass Hash Table */
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	return []int{}
}

// @lc code=end
