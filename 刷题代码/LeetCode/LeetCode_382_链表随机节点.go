package main

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {

	return Solution{head}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {

	return this.get(1)[0]
}

func (this *Solution) get(k int) []int{
	arr:=make([]int,0)
	cur:=this.head
	index:=0
	for index<k  {
		arr = append(arr,cur.Val)
		cur = cur.Next
		index++
	}
	for cur!=nil {
		x:=rand.Intn(index+1)	// 注意这个 这样才表示在 [0,index]中取一个值
								
		if x<k{
			arr[x] = cur.Val
		}
		cur = cur.Next
		index++
	}
	return arr
}
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
	
}
/*
	总结
	1. 这题刚开始做的时候理解错了，以为是每个节点出现的概率是相等的，然后就直接用个循环链表一直输出cur
       之后错了之后，看了题目，还以为是数值出现概率相等，然后搞不定就看答案了..
    2. 答案说这题用到了蓄水池算法 0.0，我也学着AC了，
    3. 之后也用自己的算法AC了，直接随机一个索引，然后去链表找到这个索引就可以了，输出这个索引对应的节点的值就可以了
    4. 之后去看博客，也自己写出了通用的蓄水池算法
*/