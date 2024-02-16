package problem115

/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */

// @lc code=start
func numDistinct(s string, t string) int {
	cache := [][]int{}
	for i := 0; i < len(s); i++ {
		arr := make([]int, len(t))
		for j := 0; j < len(t); j++ {
			arr[j] = -1
		}
		cache = append(cache, arr)
	}
	return findDistinct(s, t, 0, 0, cache)
}

func findDistinct(s, t string, sIndex, tIndex int, cache [][]int) int {
	if tIndex == len(t) {
		return 1
	}

	if sIndex == len(s) {
		return 0
	}

	if cache[sIndex][tIndex] != -1 {
		return cache[sIndex][tIndex]

	}

	count := 0
	for i := sIndex; i < len(s); i++ {
		if s[i] == t[tIndex] {
			count += findDistinct(s, t, i+1, tIndex+1, cache)
		}
	}

	cache[sIndex][tIndex] = count
	return count
}

// @lc code=end
