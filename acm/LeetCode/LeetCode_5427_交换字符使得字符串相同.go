package main

func minimumSwap(s1 string, s2 string) int {
	x2y,y2x:=0,0
	for i:=0;i<len(s1);i++{
		if s1[i]!=s2[i]{
			if s1[i]=='x'{
				x2y++
			}else{
				y2x++
			}
		}
	}
	ans:=x2y/2 + y2x/2
	x2y%=2
	y2x%=2

	if x2y==y2x{
		return ans +  x2y  + y2x
	}else{
		return -1
	}
}

/*
	总结
	1. 这题刚做的时候没有思路..想了好久也没做出来
	2. 比赛结束后看官方题解才AC了
	3. 思路是 从第一个字符串算（从第二个也可以），看看他有多少个位置需要
       把x换y，以及多少个位置要y换x，根据这个进行判断..

*/


