package main

import (
	"fmt"
)
var mp map[uint8]int

func main(){
	fmt.Println(balancedString("QEWRWWWW"))
	//fmt.Println(lengthOfLongestSubstring("QQQE"))//
	//fmt.Println(lengthOfLongestSubstring("QQQQ"))/
	//fmt.Println(lengthOfLongestSubstring("EEEE"))
}

func balancedString(s string) int {
	mp = make(map[uint8]int)
	mp['Q']=0
	mp['W']=1
	mp['E']=2
	mp['R']=3
	target:=[4]int{}
	for i:=0;i<len(s);i++{
		target[mp[s[i]]]++
	}
	balanceEle:=len(s)/4

	for i:=0;i<len(target);i++{
		target[i] = target[i]-balanceEle
		if target[i]<0{
			target[i] = 0
		}
	}
	flag:=0
	for i:=0;i<len(target);i++{
		if target[i]!=0{
			flag = 1
			break
		}
	}
	if flag==0{
		return 0
	}

	node:=Node{s:s,target:target,now:[4]int{}}
	return solve(&node)

}

type Window interface {
	Len() int				// 总窗口长度
	Add(i int)
	Sub(i int)
	ToMove() bool
	IsMax() bool
}
type Node struct {
	s string
	target [4]int
	now [4]int
}

func (n *Node) IsMax() bool{
	return false
}
func (n *Node) Len() int{
	return len(n.s)
}

func (n *Node) ToMove() bool{
	for i:=0;i<4;i++{
		if  n.now[i]<n.target[i]{
			return false
		}
	}
	return true
}

func (n *Node) Add(i int){
	n.now[mp[n.s[i]]]++
}

func (n *Node) Sub(i int){
	n.now[mp[n.s[i]]]--
}


func solve(window Window) int{
	l,r:=0,0
	inf := 1000000000000000000
	ans:=inf
	if window.IsMax(){
		ans=0
	}

	for l<=r && r<=window.Len()-1 {

		window.Add(r)
		for l<=r && window.ToMove(){
			if !window.IsMax(){
				ans = min(ans,r-l+1)	// 这个接口的ans放进循环了 0.0
			}
			window.Sub(l)
			l++
		}
		if window.IsMax(){
			ans = max(ans,r-l+1)	// 这个接口的ans放进循环了 0.0
		}
		r++
	}
	return ans
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