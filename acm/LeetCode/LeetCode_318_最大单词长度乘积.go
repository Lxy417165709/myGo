package main

import "fmt"
func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}
func maxProduct(words []string) int {
	slice:=make([]int,0)
	for i:=0;i< len(words);i++  {

		sum:=0
		for t:=0;t< len(words[i]);t++  {
			sum = sum | 1<<(words[i][t]-'a')
		}
		slice = append(slice, sum)
	}
	mx:=0
	for i:=0;i< len(slice);i++  {
		for t:=i+1;t<len(slice) ;t++  {
			if slice[i]&slice[t]==0 {
				mx = max(mx, len(words[i])*len(words[t]))
			}
		}
	}
	return mx

}
func main() {
	fmt.Println(1<<0)
	fmt.Println(maxProduct([]string{

	}))
}
/*
	总结
	1. 刚开始想到了O(xn^2)的算法，利用的是哈希+暴力遍历，x为字符串长度，x时间被用于判断是否存在相同字母(1个1个的找)
	2. 看了官方题解后，它的思路和我差不多，只不过是用位运算代替了字母相同的判别，把x时间化为了1
       比如 abc 可以表示为  00111,cde表示为 11100，两者通过&运算，如果等于0表示没有相同字母 0.0..
*/