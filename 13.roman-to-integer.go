/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package main

// @lc code=start
func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	last := 0

	for i := len(s) - 1; i >= 0; i-- {
		current := roman[s[i]]

		if current < last {
			total -= current
		} else {
			total += current
		}
		last = current
	}

	return total
}

// @lc code=end
