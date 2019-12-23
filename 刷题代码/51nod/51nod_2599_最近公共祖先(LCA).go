package _1nod

import (
	"container/list"
	"fmt"
)
var n = 0
type Node struct {
	father int
	depth int
}
var node = [1050]Node{}
var mark = [1050]int{}
var m = [1050][1050]int{}

func setDepth(){
		mylist :=list.New()
		mylist.PushBack(1)
		node[1].depth = 0
		node[1].father = 1
		for ;mylist.Len()>0;{
			top:=mylist.Front().Value.(int)
			//fmt.Println(top)
			mylist.Remove(mylist.Front())
			mark[top] = 1
			for i:=1; i<=n;i++  {
				if m[top][i] == 1 {
					//fmt.Println(i)
					if mark[i] != 1 {
						node[i].depth = node[top].depth + 1
						node[i].father = top
						mylist.PushBack(i)
					}
				}
			}
		}
}

func LCA(u,v int) int{
	uD,vD := node[u].depth,node[v].depth
	if uD >  vD{
		u,v = v,u
	}
	delta := node[v].depth - node[u].depth

	for i:=1;i<=delta;i++{
		v = node[v].father
	}
	for ;u!=v;{
		u = node[u].father
		v = node[v].father
	}
	return u

}

func main() {

	fmt.Scan(&n)
	for i:=1;i<n;i++{
		u,v:=0,0
		fmt.Scan(&u,&v)
		m[u][v] = 1
		m[v][u] = 1
	}
	setDepth()
	Q:=0
	fmt.Scan(&Q)
	for i:=1;i<=Q;i++{
		u,v:=0,0
		fmt.Scan(&u,&v)
		fmt.Println(LCA(u,v))
	}
}


/*
	总结
	1. go语言的队列可以用List实现

*/