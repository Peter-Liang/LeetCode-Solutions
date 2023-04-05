package problem_93

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func Test0(t *testing.T) {
	result := restoreIpAddresses("0000")

	assert.Equal(t, []string{"0.0.0.0"}, result)
}

func Test1(t *testing.T) {
	result := restoreIpAddresses("25525511135")

	t.Log(result)
	expected := []string{"255.255.11.135", "255.255.111.35"}

	assert.Equal(t, len(expected), len(result))
	for _, val := range result {
		assert.Contains(t, expected, val)
	}
	// assert.Equal(t, expected, result)
}

func Test2(t *testing.T) {
	result := restoreIpAddresses("101023")

	assert.Equal(t, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, result)
}

// @lc code=end
