package problem_91

import (
	"strconv"
)

var cache = make(map[string]int)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	if val, ok := cache[s]; ok {
		return val
	}

	num, _ := strconv.ParseInt(s[:2], 10, 64)
	if num > 26 {
		if num%10 == 0 {
			cache[s] = 0
			return 0
		}

		cache[s] = numDecodings(s[1:])
		return cache[s]
	}

	if num%10 == 0 {
		cache[s] = numDecodings(s[2:])
		return cache[s]
	}

	cache[s] = numDecodings(s[1:]) + numDecodings(s[2:])
	return cache[s]
}
