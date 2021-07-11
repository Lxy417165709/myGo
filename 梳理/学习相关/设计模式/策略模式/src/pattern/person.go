package pattern

import "fmt"

type Person struct {
	EatWay Strategy
}

func (p *Person) SetEatWay(s Strategy) {
	p.EatWay = s
}
func (p Person) EatRice() {
	if p.EatWay == nil{
		fmt.Println("吃饭是什么技能？我不会阿")
		return
	}
	p.EatWay.Eat()
}
