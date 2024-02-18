package problem119

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	p := make([]int, 0, rowIndex+1)
	p = append(p, 1)
	c := make([]int, 0, rowIndex+1)
	for i := 1; i <= rowIndex; i++ {
		c = c[:i+1]
		c[0] = 1
		for j := 1; j < len(p); j++ {
			c[j] = p[j-1] + p[j]
		}
		c[len(c)-1] = 1
		p, c = c, p
	}

	return p
}

// @lc code=end
