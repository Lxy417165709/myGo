package main

import "fmt"

/*
	一个仅有ab的字符串，要求a和b需要成对出现，否则不合法,即 (aa|bb)* 正则的匹配。我们可以用dfa来做这个题。
*/

func isMatch(s string) bool {
	// 状态机矩阵
	matrixOfDFA := [][]int{
		{0, 0},
		{2, 3},
		{1, 4},
		{4, 1},
		{4, 4},
	}
	// 操作
	toStateAction := map[uint8]int{'a': 0, 'b': 1}

	// 当前状态
	nowState := 1
	for i := 0; i < len(s); i++ {
		nowState = matrixOfDFA[nowState][toStateAction[s[i]]]
		fmt.Println("当前状态:", nowState, "  ", "当前判断的字符串:", s[:i+1])
	}
	return nowState == 1
}

func main() {
	s := ""
	for true {
		fmt.Scan(&s)
		fmt.Println(isMatch(s))
	}
}

/*
	测试结果:
		(1) 案例
			输入: aabbaaabb

				当前状态: 2    当前判断的字符串: a
				当前状态: 1    当前判断的字符串: aa
				当前状态: 3    当前判断的字符串: aab
				当前状态: 1    当前判断的字符串: aabb
				当前状态: 2    当前判断的字符串: aabba
				当前状态: 1    当前判断的字符串: aabbaa
				当前状态: 2    当前判断的字符串: aabbaaa
				当前状态: 4    当前判断的字符串: aabbaaab
				当前状态: 4    当前判断的字符串: aabbaaabb

			结果: false

		(2) 案例
			输入: aaaaaaaabbbbbbbb

				当前状态: 2    当前判断的字符串: a
				当前状态: 1    当前判断的字符串: aa
				当前状态: 2    当前判断的字符串: aaa
				当前状态: 1    当前判断的字符串: aaaa
				当前状态: 2    当前判断的字符串: aaaaa
				当前状态: 1    当前判断的字符串: aaaaaa
				当前状态: 2    当前判断的字符串: aaaaaaa
				当前状态: 1    当前判断的字符串: aaaaaaaa
				当前状态: 3    当前判断的字符串: aaaaaaaab
				当前状态: 1    当前判断的字符串: aaaaaaaabb
				当前状态: 3    当前判断的字符串: aaaaaaaabbb
				当前状态: 1    当前判断的字符串: aaaaaaaabbbb
				当前状态: 3    当前判断的字符串: aaaaaaaabbbbb
				当前状态: 1    当前判断的字符串: aaaaaaaabbbbbb
				当前状态: 3    当前判断的字符串: aaaaaaaabbbbbbb
				当前状态: 1    当前判断的字符串: aaaaaaaabbbbbbbb

			输出: true
*/
