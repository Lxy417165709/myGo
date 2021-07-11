package 双指针

import (
	"fmt"
	"strconv"
	"strings"
)

func compressString(S string) string {
	first, last := 0, 0
	result := ""
	for last < len(S) {
		for last+1 < len(S) && S[last+1] == S[last] {
			last++
		}

		result += string(S[first]) + fmt.Sprintf("%d", last-first+1) // 可以采用stringBuilder优化
		first = last + 1
		last = first
	}
	if len(result) < len(S) {
		return result
	}
	return S
}
// 通过	80 ms	7.9 MB	Golang

// 优化 ()
func compressString(S string) string {
	first, last := 0, 0
	result := strings.Builder{}
	for last < len(S) {
		for last+1 < len(S) && S[last+1] == S[last] {
			last++
		}
		result.WriteByte(S[first])
		//result.WriteString(fmt.Sprintf("%d", last-first+1))
		result.WriteString(strconv.Itoa(last - first + 1))
		first = last + 1
		last = first
	}
	if result.Len() < len(S) {
		return result.String()
	}
	return S
}
// 	通过	4 ms	3.4 MB	Golang

// 再优化
func compressString(S string) string {
	first, last := 0, 0
	result := strings.Builder{}
	for last < len(S) {
		for last+1 < len(S) && S[last+1] == S[last] {
			last++
		}
		result.WriteByte(S[first])
		result.WriteString(strconv.Itoa(last - first + 1)) // 优化了这里
		first = last + 1
		last = first
	}
	if result.Len() < len(S) {
		return result.String()
	}
	return S
}
// 	通过	0 ms	3.1 MB	Golang

/*
	链接： https://leetcode-cn.com/problems/compress-string-lcci/
*/
