package pattern

import "fmt"

type Strategy2 struct {
}
func (s Strategy2)Eat() {
	fmt.Println("坐着吃饭！")
}
