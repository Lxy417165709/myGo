package pattern

import "fmt"

type Strategy1 struct {
}
func (s Strategy1)Eat() {
	fmt.Println("站着吃饭！")
}
