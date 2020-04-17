package main

import (
	"fmt"
	"pattern"
)

func main() {
	mgb := pattern.NewMagicButton(pattern.TV{},
		pattern.Light{},
		pattern.WaterHeater{},
	)

	fmt.Println("我到店了~~~我按下魔术按钮后(打开)~~~")
	mgb.On()
	fmt.Println()
	fmt.Println("店要关门了~~~我按下魔术按钮后(关闭)~~~")
	mgb.Off()

	// 我到店了~~~我按下魔术按钮后(打开)~~~
	// 电视机打开了
	// 电灯打开了
	// 热水机启动啦
	//
	// 店要关门了~~~我按下魔术按钮后(关闭)~~~
	// 电视机关闭了
	// 电灯关闭了
	// 热水机关闭了
}
// 外观模式就是为了简化用户对多个实体的操作而出现的设计模式
