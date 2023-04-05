package problem_93

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	splits := getSplits(s, 4)
	for _, v := range splits {
		result = append(result, strings.Join(v, "."))
	}

	return result
}

func getSplits(s string, parts int) [][]string {
	if len(s) == 0 {
		return nil
	}

	if parts == 1 {
		if len(s) > 1 && s[0] == '0' {
			return nil
		}

		if len(s) > 3 {
			return nil
		}

		val, _ := strconv.ParseInt(s, 10, 64)
		if val > 255 {
			return nil
		}

		return [][]string{{s}}
	}

	result := make([][]string, 0)
	for i := 1; i <= 4; i++ {
		if i > len(s) {
			break
		}

		ip := s[0:i]

		val, _ := strconv.ParseInt(ip, 10, 64)
		if val > 255 {
			break
		}

		sub := getSplits(s[i:], parts-1)
		if sub != nil {
			for _, v := range sub {
				res := append([]string{ip}, v...)
				result = append(result, res)
			}
		}

		if ip == "0" {
			break
		}
	}

	return result
}

// @lc code=end
