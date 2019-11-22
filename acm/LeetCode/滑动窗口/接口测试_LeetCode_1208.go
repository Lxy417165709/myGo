package main

func equalSubstring(s string, t string, maxCost int) int {

	model:=Node{s,t,maxCost}
	return solve(model)
}

type Window interface {
	Len() int				// 总窗口长度

	ToMove(nowValue int) bool
	Add(nowValue,i int) int
	Sub(nowValue,i int) int
}
type Node struct {
	s string
	t string
	maxCost int
}

func (n Node) GetPointValue(i int) int{
	return abs(int(n.s[i])-int(n.t[i]))
}

func (n Node) Len() int{
	return len(n.t)
}

func (n Node) ToMove(nowValue int) bool{
	return nowValue>n.maxCost
}

func (n Node) Add(nowValue,i int) int{
	nowValue = nowValue + n.GetPointValue(i)
	return nowValue
}

func (n Node) Sub(nowValue,i int) int{
	nowValue = nowValue - n.GetPointValue(i)
	return nowValue
}

// 解决了1208的滑动窗口题
func solve(window Window) int{
	l,r:=0,0
	ans:=0

	nowValue:=0
	for l<=r && r<=window.Len()-1 {

		// 这里可以抽象
		nowValue = window.Add(nowValue, r)
		for l<=r && window.ToMove(nowValue){
			nowValue = window.Sub(nowValue,l)
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
func abs(a int) int{
	if a>0{
		return int(a)
	}
	return int(-a)
}