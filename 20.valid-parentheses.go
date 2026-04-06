/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package main

// @lc code=start
func isValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if opening, ok := pairs[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

// @lc code=end
