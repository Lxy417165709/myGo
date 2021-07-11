package 杂物

import (
	"fmt"
)

var m [X][Y]int
var mark[X][Y]bool
var dx = []int{0,-1,0,1}
var dy = []int{1,0,-1,0}

//var dx = []int{-1,-1,1,1,-2,-2,2,2}
//var dy = []int{2,-2,2,-2,1,-1,1,-1}
const(
	X = 5
	Y = 10
)
type Node struct {
	x int
	y int
	//lay int
}

func judge(node Node) bool{
	if node.x<0 || node.x>=X || node.y<0 || node.y>=Y {
		return false
	}
	if mark[node.x][node.y]==true {
		return false
	}
	return true
}

func markdown(node Node) {
	mark[node.x][node.y]=true
}
func showMap(){
	fmt.Println("-----------------------")
	for i:=0;i<X ;i++  {
		for t:=0;t<Y ;t++  {
			fmt.Printf("%d\t",m[i][t])
		}
		fmt.Println()
	}
	fmt.Println("-----------------------")
}

// 遍历方式1
func search1(){
	m = [X][Y]int{}
	mark = [X][Y]bool{}

	queue:=make([]Node,0)
	beginNode := Node{X/2,Y/2}
	queue = append(queue, beginNode)
	markdown(beginNode)

	step:=0

	for len(queue)!=0{
		top := queue[0]
		queue = queue[1:]

		m[top.x][top.y] = step
		step++
		for i:=0;i< len(dx);i++  {
			newNode:=Node{top.x + dx[i],top.y + dy[i]}
			if	judge(newNode){
				queue = append(queue,newNode)
				markdown(newNode)
			}
		}
	}

}

// 遍历方式1_1 --- 加标记字段时
//func search1_1(){
//	m = [X][Y]int{}
//	mark = [X][Y]bool{}
//
//	queue:=make([]Node,0)
//	beginNode := Node{X/2,Y/2,0}
//	queue = append(queue, beginNode)
//	markdown(beginNode)
//
//
//	for len(queue)!=0{
//		top := queue[0]
//		queue = queue[1:]
//		m[top.x][top.y] = top.lay
//		for i:=0;i< len(dx);i++  {
//			newNode:=Node{top.x + dx[i],top.y + dy[i],top.lay+1}
//			if	judge(newNode){
//				queue = append(queue,newNode)
//				markdown(newNode)
//			}
//		}
//	}
//
//}
//遍历方式2
func search2(){
	m = [X][Y]int{}
	mark = [X][Y]bool{}

	queue:=make([]Node,0)
	beginNode := Node{X/2,Y/2}
	queue = append(queue, beginNode)
	markdown(beginNode)

	step:=0
	for len(queue)!=0{
		size:=len(queue)
		for k:=0;k<size ;k++  {
			top := queue[0]
			queue = queue[1:]
			m[top.x][top.y] = step
			//step++	// 这可以写在循环外，那就表示循环的第几层


			for i:=0;i< len(dx);i++  {
				newNode:=Node{top.x + dx[i],top.y + dy[i]}
				if	judge(newNode){
					queue = append(queue,newNode)
					markdown(newNode)
				}
			}
		}
		step++
	}

}

//遍历方式3
func search3(){
	m = [X][Y]int{}
	mark = [X][Y]bool{}

	queue:=make([]Node,0)
	beginNode := Node{X/2,Y/2}
	queue = append(queue, beginNode)
	markdown(beginNode)

	step:=0
	for len(queue)!=0{
		top := queue[0]
		queue = queue[1:]
		m[top.x][top.y] = step
		tmp:=make([]Node,0)
		for i:=0;i< len(dx);i++  {
			newNode:=Node{top.x + dx[i],top.y + dy[i]}
			if	judge(newNode){
				tmp = append(tmp, newNode)
				//queue = append(queue,newNode)
				markdown(newNode)
			}
		}
		//for i:=0;i< len(tmp)/2;i++  {
		//	tmp[i],tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i],tmp[i]
		//} 遍历方式 3-1
		queue = append(tmp,queue...)
		step++

	}

}

func main() {
	search3()
	showMap()
}
/*
	遍历方式1
	-----------------------
	48	44	37	28	18	8	16	26	35	42
	45	38	29	19	9	2	6	14	24	33
	39	30	20	10	3	0	1	5	13	23
	46	40	31	21	11	4	7	15	25	34
	49	47	41	32	22	12	17	27	36	43
	-----------------------


	遍历方式2
	-----------------------
	48	44	37	28	18	8	16	26	35	42
	45	38	29	19	9	2	6	14	24	33
	39	30	20	10	3	0	1	5	13	23
	46	40	31	21	11	4	7	15	25	34
	49	47	41	32	22	12	17	27	36	43
	-----------------------
	2种是一样的，0.0..


	遍历方式3(DFS)
	-----------------------
	15	14	13	12	11	10	9	8	7	6
	16	36	37	38	39	47	45	43	41	5
	17	18	19	20	48	0	1	2	3	4
	35	34	33	21	22	49	46	44	42	40
	32	31	30	29	23	24	25	26	27	28
	-----------------------

    遍历方式3-1(DFS) 其实就是逆序了
	-----------------------
	11	12	13	38	19	20	21	37	33	32
	10	39	14	17	18	48	22	23	36	31
	9	40	15	16	47	0	49	24	35	30
	8	41	42	43	45	1	46	25	34	29
	7	6	5	4	3	2	44	26	27	28
	-----------------------




	总结
	1. 其实遍历方式2和遍历方式1遍历顺序是一样的，只不过遍历方式2可以通过for循环分层，而遍历方式1的分层需要
	   在Node结构体里加层的标记字段lay
	2. 遍历方式3是一路走到底的，
	3. 总的来说，遍历的顺序和元素的取出顺序有关(取出顺序与数据结构有关，可以使用栈，普通队列，优先队列......)，



*/