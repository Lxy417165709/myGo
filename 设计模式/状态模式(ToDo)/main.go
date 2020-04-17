package main

import (
	"fmt"
	"pattern"
)

func main() {
	me := pattern.GetMyself()	// 获取自己 (初始化为开心状态)
	fmt.Println("初始化状态:",me.ToString())

	me.GoToSendTakeOut()
	fmt.Println("我去送外卖后，状态变化为:", me.ToString())

	me.BeComplained()
	fmt.Println("我被批评后，状态变化为:",me.ToString())

	me.DrinkCaoMeiYiJunDuo()
	fmt.Println("我喝了草莓益菌多后，状态变化为:",me.ToString())

	me.BeComplained()
	fmt.Println("我被批评后，状态变化为:",me.ToString())

	me.GoToSendTakeOut()
	fmt.Println("我去送外卖后，状态变化为:", me.ToString())
}
