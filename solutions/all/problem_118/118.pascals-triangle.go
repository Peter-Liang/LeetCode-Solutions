package problem118

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		p := res[i-1]
		c := []int{1}
		for j := 1; j < len(p); j++ {
			c = append(c, p[j-1]+p[j])
		}
		c = append(c, 1)
		res[i] = c
	}

	return res
}

// @lc code=end
