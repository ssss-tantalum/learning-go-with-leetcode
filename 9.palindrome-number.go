/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package main

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}

	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}

// @lc code=end
