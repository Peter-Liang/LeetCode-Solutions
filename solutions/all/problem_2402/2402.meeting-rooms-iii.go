package problem2402

import (
	"math"
	"slices"
)

/*
 * @lc app=leetcode id=2402 lang=golang
 *
 * [2402] Meeting Rooms III
 */

// @lc code=start
func mostBooked(n int, meetings [][]int) int {
	bookCount := make([]int, n)
	booked := make([][]int, n)
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})
	for _, m := range meetings {
		earliest, idx := math.MaxInt, 0
		found := false
		for i, v := range booked {
			if v == nil || v[1] <= m[0] {
				booked[i] = m
				bookCount[i]++
				found = true
				break
			} else if v[1] < earliest {
				earliest = v[1]
				idx = i
			}
		}

		if !found {
			duration := m[1] - m[0]
			booked[idx][0] = booked[idx][1]
			booked[idx][1] = booked[idx][0] + duration
			bookCount[idx]++
		}
	}

	max, idx := 0, 0
	for i, v := range bookCount {
		if v > max {
			max = v
			idx = i
		}
	}

	return idx
}

// @lc code=end
