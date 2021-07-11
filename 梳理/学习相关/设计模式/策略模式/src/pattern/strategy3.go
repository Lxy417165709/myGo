package pattern

import "fmt"

type Strategy3 struct {
}
func (s Strategy3)Eat() {
	fmt.Println("倒立着吃饭！")
}
