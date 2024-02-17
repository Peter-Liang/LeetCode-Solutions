package problem1642

import (
	"slices"
)

/*
 * @lc app=leetcode id=1642 lang=golang
 *
 * [1642] Furthest Building You Can Reach
 */

// @lc code=start
func furthestBuilding(heights []int, bricks int, ladders int) int {
	if len(heights) < 2 {
		return 0
	}

	diffs := make([]int, 0)
	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			continue
		}

		useBrick := true
		if bricks < diff {
			if ladders >= len(heights)-i {
				return len(heights) - 1
			}

			if ladders == 0 {
				return i - 1
			}

			slices.Sort(diffs)
			ladders--
			useBrick = false
			if len(diffs) > 0 && diff < diffs[len(diffs)-1] {
				bricks += diffs[len(diffs)-1]
				diffs = diffs[:len(diffs)-1]
				useBrick = true
			}
		}

		if useBrick {
			bricks -= diff
			if diff > 1 {
				diffs = append(diffs, diff)
			}
		}
	}

	return len(heights) - 1
}

// @lc code=end
