package main

func sortedListToBST(head *ListNode) *TreeNode {
	return solve(head)
}

func solve(head *ListNode) *TreeNode{

	if head==nil{
		return nil
	}

	var pre *ListNode = nil
	fast,slow:=head,head	//fast = head or fast = head.Next 都可以

	// 优美了一点
	for fast!=nil && fast.Next!=nil{
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}



	var left *TreeNode = nil
	// 第二次虽然知道要把pre.Next置空，但是没有判断pre==nil时，其实左子树就不用递归了，
	// 于是导致左子树无穷...
	// 之前写法是
	// root := &TreeNode{
	//		Val: slow.Val,
	//		Left: solve(head),
	//		Right: solve(slow.Next),
	//	}
	if pre!=nil{
		pre.Next = nil
		left=solve(head)
	}else{
		left = nil
	}
	root := &TreeNode{
		Val: slow.Val,
		Left: left,
		Right: solve(slow.Next),
	}

	return root
}

/*
	总结
	1. 刚刚开始想到了，先把链表信息存储在数组之中，之后按照108的思路就可以AC了，但是觉得很简单，
		所以要想其他的。
	2. 然后想到了平衡二叉树左偏右偏..但是不知道怎么写
	3. 看了官方题解之后，他也是求中点，就求链表中点，之后递归就可以了。
		..很无语的就是自己还因为几个条件调试了半天..
	4. 官方有我想要的解法 有序链表转换二叉搜索树（底-顶递归法，复杂度O(N)）这个，去实现！

*/
