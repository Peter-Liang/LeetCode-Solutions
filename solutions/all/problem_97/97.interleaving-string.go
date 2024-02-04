package problem_97

/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[int]map[int]bool)
	var check func(int, int, int) bool
	check = func(i, j, k int) bool {
		if len(s3) <= k {
			return true
		}

		if dp[i][j] {
			return false
		}

		if i < len(s1) && s1[i] == s3[k] {
			if check(i+1, j, k+1) {
				return true
			}

			if dp[i+1] == nil {
				dp[i+1] = make(map[int]bool)
			}
			dp[i+1][j] = true
		}

		if j < len(s2) && s2[j] == s3[k] {
			if check(i, j+1, k+1) {
				return true
			}
			if dp[i] == nil {
				dp[i] = make(map[int]bool)
			}
			dp[i][j+1] = true
		}

		if dp[i] == nil {
			dp[i] = make(map[int]bool)
		}
		dp[i][j] = true
		return false
	}

	return check(0, 0, 0)
}

// @lc code=end
