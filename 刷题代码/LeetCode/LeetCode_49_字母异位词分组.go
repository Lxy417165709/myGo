package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string]int)
	sli:=make([][]string,0)

	tmp:=make([]byte,0)
	for i:=0;i<len(strs);i++{

		tmp=[]byte(strs[i])
		sort.Slice(tmp,func (p,t int) bool{
			return tmp[p]<tmp[t]
		})
		if x,ok:=mp[string(tmp)];!ok{
			sli = append(sli,[]string{strs[i]})
			mp[string(tmp)] = len(sli)-1
		}else{
			sli[x] = append(sli[x],strs[i])
		}

	}
	return sli
}


func main() {

}

/*
	总结
	1. 这题我是对每个字符串排序后，然后看是否相等来判断他们是否是同一组的 0.0
	2. 官方也有使用字符计次的，但是我觉得那个用质数分解的大佬很牛逼！
		如下


算术基本定理，又称为正整数的唯一分解定理，即：每个大于1的自然数，要么本身就是质数，要么可以写为2个以上的质数的积，而且这些质因子按大小排列之后，写法仅有一种方式。
利用这个，我们把每个字符串都映射到一个正数上。

用一个数组存储质数 prime = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}。
这样，我们把'a'对应2,'b'对应3，以此类推
于是异位词之间对应的数是相等的0.0
作者：windliang
链接：https://leetcode-cn.com/problems/group-anagrams/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--16/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/