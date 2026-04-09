/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package main

// @lc code=start
func romanToInt(s string) int {
	// ローマ数字と整数の対応表をマップにする
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	length := len(s)

	for i := 0; i < length; i++ {
		// 現在の文字の値を取得
		current := romanMap[s[i]]

		// 「次の文字」が存在し、かつ「現在の値 < 次の値」の場合
		if i+1 < length && current < romanMap[s[i+1]] {
			// 今の値を引く
			result -= current
		} else {
			// 今の値を足す
			result += current
		}
	}

	return result
}

// @lc code=end
