package problem1481

import "slices"

/*
 * @lc app=leetcode id=1481 lang=golang
 *
 * [1481] Least Number of Unique Integers after K Removals
 */

// @lc code=start
func findLeastNumOfUniqueInts(arr []int, k int) int {
	if k >= len(arr) {
		return 0
	}

	counter := make(map[int]int, len(arr))
	for _, v := range arr {
		counter[v]++
	}

	sortedCount := make([]int, len(counter))
	i := 0
	for _, v := range counter {
		sortedCount[i] = v
		i++
	}
	slices.Sort(sortedCount)
	for i, v := range sortedCount {
		k -= v
		if k < 0 {
			return len(sortedCount) - i
		}
	}

	return 0
}

// @lc code=end
