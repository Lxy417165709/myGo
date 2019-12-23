package main



type Leaderboard struct {
	goal map[int]int
	idToGoal map[int]int
}

func Constructor() Leaderboard {
	return Leaderboard{make(map[int]int),make(map[int]int)}
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
	if x,ok:=this.idToGoal[playerId];ok{
		this.idToGoal[playerId] = x + score
		this.goal[x]--
		this.goal[x+score]++
	}else{
		this.goal[score]++
		this.idToGoal[playerId] = score
	}
}

func (this *Leaderboard) Top(K int) int {
	ans:=0
	for i:=100000;i>=0 && K>0;i--{
		if this.goal[i]!=0{
			if this.goal[i]>=K{
				ans+=i*K
				K = K - K
			}else{
				ans+=i*this.goal[i]
				K = K - this.goal[i]
			}
		}

	}
	return  ans
}

func (this *Leaderboard) Reset(playerId int)  {
	num:=this.idToGoal[playerId]
	this.goal[num]--
	delete(this.idToGoal,playerId)
}
func main() {

}
