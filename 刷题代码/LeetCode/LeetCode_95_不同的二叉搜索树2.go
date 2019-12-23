package main

var mp map[int][]*TreeNode
var offset uint8= 20
func generateTrees(n int) []*TreeNode {
	mp = make(map[int][]*TreeNode)
	if n==0{
		return []*TreeNode{}
	}
	ans := solve(1,n)

	return ans
}

func solve(l,r int) []*TreeNode{
	if l>r{
		return []*TreeNode{
			nil,
		}
	}

	if l==r{
		return []*TreeNode{
			&TreeNode{l,nil,nil},
		}
	}
	if x,ok:=mp[l<<offset|r];ok{
		return x
	}
	ans:=make([]*TreeNode,0)

	for i:=l;i<=r;i++{

		left:=solve(l,i-1)
		right:=solve(i+1,r)

		for t:=0;t<len(left);t++{
			for j:=0;j<len(right);j++{
				root:=&TreeNode{i,left[t],right[j]}
				ans = append(ans,root)
			}
		}

	}
	mp[l<<offset|r] = ans
	return mp[l<<offset|r]
}
func main() {

}
/*
	总结
	1. 这题目我用分治做出来了 0.0,之前是求个数，这次是求具体的树
	2. n==0的时候给坑了一把...我觉得应该就是返回[[]] 0.0..
	3. 其实我写的还有存在重叠子问题 0.0 我优化一下
	4. 我采用了记忆化搜索优化了下，然后发现内存消耗更少了，但是时间只提升了4ms 0.0
*/
