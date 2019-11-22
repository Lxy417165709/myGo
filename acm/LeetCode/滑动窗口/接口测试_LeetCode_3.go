package main

func lengthOfLongestSubstring(s string) int {
	node:=Node{s,make(map[byte]int),-1}
	return solve(&node)
}

type Window interface {
	Len() int				// 总窗口长度
	Add(i int)
	Sub(i int)
	ToMove() bool
}
type Node struct {
	s string
	mp map[byte]int
	mark int
}


func (n *Node) Len() int{
	return len(n.s)
}

func (n *Node) ToMove() bool{

	return n.mark != -1
}

func (n *Node) Add(i int){
	n.mp[n.s[i]]++
	if n.mp[n.s[i]]==2{
		n.mark = i
	}
}

func (n *Node) Sub(i int){
	n.mp[n.s[i]]--

	if n.mark!=-1 && n.s[i]==n.s[n.mark]{
		n.mark = -1
	}

}

// 解决了3的滑动窗口题
func solve(window Window) int{
	l,r:=0,0
	ans:=0

	for l<=r && r<=window.Len()-1 {

		// 这里可以抽象
		window.Add(r)
		for l<=r && window.ToMove(){
			window.Sub(l)
			l++
		}

		ans = max(ans,r-l+1)
		r++
	}
	return ans
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}