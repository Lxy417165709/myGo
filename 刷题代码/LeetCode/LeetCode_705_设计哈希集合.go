package main


type Node struct{
	Val int
	Next *Node
}

type List struct{
	head *Node
	end *Node
}

func (this *List)listAdd(key int) {
	if this.head == nil{
		this.head = &Node{key,nil}
		this.end = this.head
	}else{
		this.end.Next = &Node{key,nil}
		this.end = this.end.Next
	}
}

func (this *List)listSearch(key int) bool{
	tmp := this.head
	for tmp!=nil{
		if tmp.Val == key{
			return true
		}
		tmp = tmp.Next
	}
	return false
}


func (this *List)listRemove(key int) {

	if this.head == nil{
		return
	}
	for this.head.Val == key{
		this.head = this.head.Next
		if this.head==nil{
			this.end = nil
			return
		}
	}
	tmp := this.head
	for tmp.Next!= nil{
		if tmp.Next.Val == key{
			if tmp.Next == this.end{
				this.end = tmp
			}
			tmp.Next = tmp.Next.Next
		}else{
			tmp = tmp.Next
		}
	}
}





type MyHashSet struct {
	hashArr [10000]List
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{[10000]List{}}
}


func (this *MyHashSet) Add(key int)  {
	this.hashArr[key%10000].listAdd(key)
}


func (this *MyHashSet) Remove(key int)  {
	this.hashArr[key%10000].listRemove(key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.hashArr[key%10000].listSearch(key)
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

 /*
 	总结
 	1. 这是使用拉链法的哈希
 	2. 代码写得有点难看..
 */
