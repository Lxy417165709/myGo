package main

import (
	"fmt"
	"sort"
)
type arr [][]int
func (a arr)Len() int{
	return len(a)
}
func (a arr)Swap(i,j int){
	a[i],a[j] = a[j],a[i]
}
func (a arr)Less(i,j int) bool{
	return a[i][0]<a[j][0]
}
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Sort(arr(slots1))
	sort.Sort(arr(slots2))
	len1:= len(slots1)
	len2:= len(slots2)
	for i,t:=0,0;i<len1 && t<len2;   {
		for i<len1 && t<len2 && slots1[i][0]>=slots2[t][1] {
			t++
		}
		for i<len1 && t<len2 && slots1[i][1]<=slots2[t][0] {
			i++
		}
		if i<len1 && t<len2 {
			a,b:=slots1[i],slots2[t]
			x:=get(a,b)
			if x>=duration {
				mx:=max(a[0],b[0])
				return []int{
					mx,mx+duration,
				}
			}
			i++	// 很重要！！ t++居然不可以..
		}


	}


	return []int{}
}
func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
func get(a []int,b[]int) int{
	tmpa:=make([]int,2)
	tmpb:=make([]int,2)
	copy(tmpa,a)
	copy(tmpb,b)
	if tmpa[0]>tmpb[0]{
		tmpa,tmpb=tmpb,tmpa
	}
	return min(tmpa[1]-tmpb[0],tmpb[1]-tmpb[0])
}


func main() {

	fmt.Println(minAvailableDuration([][]int{
		{216397070,363167701},{98730764,158208909},{441003187,466254040},{558239978,678368334},{683942980,717766451},
	},[][]int{
	{50490609,222653186},{512711631,670791418},{730229023,802410205},{812553104,891266775},{230032010,399152578},
	},456085))
}
/*
	总结
	1. 第一次做的时候，我用了暴力的方法，之后TLE了
	2. 之后想了想，在纸上模拟了一下，决定用双指针，但是由于上面的小错误（i++没写）于是也TLE
    3. 加上之后就AC了
	4. 这种题型我不太会 0.0 不太熟悉！但是我AC了哈哈哈哈哈
*/