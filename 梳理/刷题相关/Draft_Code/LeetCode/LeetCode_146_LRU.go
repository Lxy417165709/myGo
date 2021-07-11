package main

import "fmt"

type LRUCache struct {
	Cap int
	KeyToNode map[int]*MyNode
	List DoubleList
}

type MyNode struct{
	Key,Val int
	Next *MyNode
	Pre *MyNode
}


func Constructor(capacity int) LRUCache {

	return LRUCache{capacity,make(map[int]*MyNode),ConstructorOfList()}
}


func (this *LRUCache) Get(key int) int {
	node:=this.KeyToNode[key]
	if node==nil{
		return -1
	}


	this.List.rev(node)
	this.List.apd(node)

	return node.Val
}


func (this *LRUCache) Put(key int, value int)  {


	tmp:=this.KeyToNode[key]
	if tmp!=nil{
		this.List.rev(tmp)
	}


	node := &MyNode{key,value,nil,nil}
	if this.Cap==this.List.Length{
		tmp:=this.List.pop()
		this.KeyToNode[tmp.Key] = nil
	}

	this.List.apd(node)
	this.KeyToNode[key] = node
}

type DoubleList struct{
	Head,Tail *MyNode
	Length int
}
func ConstructorOfList() DoubleList{
	head:=&MyNode{-1,-1,nil,nil}
	tail:=&MyNode{-1,-1,nil,nil}
	head.Next = tail
	tail.Pre = head
	tmp:=DoubleList{head,tail,0}
	return tmp
}

func (this *DoubleList) pop() *MyNode{
	// if this.Head.Next == this.Tail{
	//     fmt.println("出错！")
	//     return
	// }

	ans:=this.Head.Next
	this.Head.Next.Next.Pre = this.Head
	this.Head.Next = this.Head.Next.Next
	this.Length--
	return ans

}

func (this *DoubleList) apd(node *MyNode){
	node.Pre = this.Tail.Pre
	node.Next = this.Tail
	this.Tail.Pre.Next = node
	this.Tail.Pre = node
	this.Length++
}

func (this *DoubleList) rev(node *MyNode){
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	this.Length--

}





/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */





func main() {
	fmt.Println("wwww")

}
/*
	总结
	1. 这题是手撕的第二题，老是出错...刚开始还理解错题意了
	2. 提交了很多次都错了，第一次是因为粗心大意，删除了双向列表的头后没有删除映射，之后错是因为题目没说清楚（如果密钥不存在，则写入其数据值，存在时它没说）
    3. 这题目我采用了双向链表法 + 哈希，在双向链表这里浪费了大量的时间，之后想到用2个节点，空头节点和空尾节点，这样可以在删除时少很多判断
       这道题 我把key映射到一个 链表节点，get时删除对应的点，并加入到双向链表的尾部，put时先看是否存在映射，是的话先删除
       之后如果长度要超过容量了,那就删除双向链表的头部，在把节点加入到最后，并设置映射关系；如果没超过，则直接加入到最后，并设置映射关系
    4. 要多多手撕鸭！！！！
*/