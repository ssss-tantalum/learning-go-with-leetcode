/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	// key: 数値, Value: インデックス
	m := make(map[int]int)

	for i, num := range nums {
		// 合計して target になるために必要な「相方」を計算
		complement := target - num
		// マップの中に「相方」がいるか
		if j, ok := m[complement]; ok {
			// いた場合、そのインデックスと現在の i を返す
			return []int{j, i}
		}
		// いなかったら、今の数値をマップに代入する
		m[num] = i
	}

	return []int{}
}

// @lc code=end
