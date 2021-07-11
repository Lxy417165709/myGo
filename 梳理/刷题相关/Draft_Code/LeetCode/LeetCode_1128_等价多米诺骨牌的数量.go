package main

import "sort"

func numEquivDominoPairs(dominoes [][]int) int {

	for i:=0;i<len(dominoes);i++{
		if dominoes[i][0]>dominoes[i][1]{
			dominoes[i][0],dominoes[i][1] = dominoes[i][1],dominoes[i][0]
		}
	}

	sort.Slice(dominoes,func (i,j int) bool{
		if dominoes[i][0] == dominoes[j][0]{
			return dominoes[i][1]<dominoes[j][1]
		}
		return dominoes[i][0]<dominoes[j][0]
	})

	l,r:=0,0
	ans:=0
	for l<=r && r<len(dominoes){
		tmp:=0
		for r<len(dominoes) && dominoes[l][0]==dominoes[r][0] && dominoes[l][1]==dominoes[r][1]{
			tmp++
			r++
		}
		if tmp&1==0{
			ans += tmp/2 *(tmp-1)
		}else{
			ans += (tmp-1)/2*tmp
		}
		l = r
	}
	return ans
}

/*
	总结
	1. 这题我是使用排序 + 双指针做的
	2. 其实也可以使用哈希做 0.0 (哈希可以映射为一个数)
*/
