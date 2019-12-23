package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == ""{
		return 0
	}

	for i:=0;i< len(haystack);i++  {
		if haystack[i] == needle[0] {
			ans:=i
			t:=0
			for ;t<len(needle) && i< len(haystack);i,t = i+1,t+1  {
				if haystack[i] != needle[t] {
					break
				}
			}
			if t==len(needle) {
				return ans
			}else{
				i = ans	// 注意这句，一定要退回  比如mismiss 寻找miss 没有这句会返回-1

			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("mismiss","miss"))
}


/*
	总结
	1. 代码有注意点，已在上面标出
	2. 这题的意思是 字符串A是否包含子串B
	3. 该题可用KMP算法解答，我这个是暴力的，时间复杂度大
	4. 该题暴力解法还能更优美！使用切片!!!
*/