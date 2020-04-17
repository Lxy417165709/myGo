package pattern

import "fmt"

type Light struct{}

func (l Light)TurnOn(){
	fmt.Println("电灯打开了")
}

func (l Light)TurnOff(){
	fmt.Println("电灯关闭了")
}
