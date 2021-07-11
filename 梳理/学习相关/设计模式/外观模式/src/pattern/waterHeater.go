package pattern

import "fmt"

type WaterHeater struct{}

func (wh WaterHeater)On(){
	fmt.Println("热水机启动啦")
}

func (wh WaterHeater)Off(){
	fmt.Println("热水机关闭了")
}
