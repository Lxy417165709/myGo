package main

func canWinNim(n int) bool {
	if n%4==0 {
		return false
	}else{
		return true
	}
}

func main() {

}

/*
	总结
	1. 该题考查博弈，博弈要注意必输态，该题必输态是n为4的倍数
*/
