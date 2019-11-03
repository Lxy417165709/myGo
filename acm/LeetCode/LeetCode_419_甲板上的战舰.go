package main

func countBattleships(board [][]byte) int {
	ans:=0
	for i:=0;i<len(board);i++{
		for t:=0;t<len(board[i]);t++{
			if board[i][t]=='X'{
				if (i==0 || board[i-1][t]=='.') &&(t==0  || board[i][t-1]=='.'){
					ans++
				}
			}
		}
	}
	return ans

}

func main() {


}


/*
	总结
	1. 由于题目设置的一些规则，我们可以知道，如果X的左边和上边不是X的话，说明
	   就多了一艘战舰 （这与我们的遍历方式有关，如果是从下到上，从右到左的话
	   ，那么该条件就是X的下边和右边不是X的话，说明多了一艘战舰	）
*/