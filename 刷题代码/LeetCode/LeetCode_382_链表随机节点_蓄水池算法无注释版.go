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

func (this *Solution) poolAlgorithm(k int) []int {
	pool := make([]int, k)
	cur := this.head
	pointer := 0
	for cur != nil {
		random := rand.Intn(pointer + 1)
		if random < k {
			pool[random] = cur.Val
		}
		pointer ++
		cur = cur.Next
	}
	return pool
}
