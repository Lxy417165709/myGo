package main

import "math/rand"
type Solution struct {
	head *ListNode
}


func Constructor(head *ListNode) Solution {

	return Solution{head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {

	return this.poolAlgorithm(1)[0]
}
// 蓄水池算法，等概选择k个元素入蓄水池
func (this *Solution) poolAlgorithm(k int) []int{
	// 蓄水池
	pool := make([]int,k)	// 创建一个大小为k的蓄水池(如果下面代码没有注释，那么初始容量应该为0)
	cur := this.head
	pointer := 0
	// 先把前 k 元素加入蓄水池 (这段代码可以省略，因为下面random在选前k个之前，概率为1)
	//for pointer < k {
	//	pool = append(pool,cur.Val)
	//	cur = cur.Next
	//	pointer++
	//}

	// 扫描链表， 每次以 pointer/k 的概率与蓄水池中的元素进行置换
	for cur != nil{
		random := rand.Intn(pointer + 1)
		if random < k {
			pool[random] = cur.Val
		}
		pointer ++
		cur = cur.Next
	}
	return pool
}