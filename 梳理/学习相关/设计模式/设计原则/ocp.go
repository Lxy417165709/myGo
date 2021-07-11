package main

import "fmt"

// 我要去送外卖啦，外卖有可能是 奶茶、小吃、口红....
// 来一段优美的ocp代码~

// 结构: 送外卖
type DeliverTakeOut struct {
}

// 送外卖执行
func (dt *DeliverTakeOut) Deliver(takeOut TakeOut) {
	takeOut.BeDelivered()
}

// 接口: 外卖
type TakeOut interface {
	BeDelivered()
}

// 实现外卖接口的结构: 奶茶
type MilkyTea struct {
}

func (mt MilkyTea) BeDelivered() {
	fmt.Println("我是奶茶，我给外卖小哥带走了")
}

// 实现外卖接口的结构: 小吃
type Snack struct {
}

func (s Snack) BeDelivered() {
	fmt.Println("我是小吃，我给外卖小哥带走了")
}

// 实现外卖接口的结构: 口红
type Lipstick struct {
}

func (lp Lipstick) BeDelivered() {
	fmt.Println("我是口红，我给外卖小哥带走了")
}

// 接下来，如果有其他可以送的
// 比如 百香思慕雪，那就可以再创建一个 百香思慕雪结构，让它实现 外卖接口
// 这样它就可以被送了
// ....

func main() {
	dto := DeliverTakeOut{}
	dto.Deliver(MilkyTea{})
	dto.Deliver(Snack{})
	dto.Deliver(Lipstick{})

	// 我是奶茶，我给外卖小哥带走了
	// 我是小吃，我给外卖小哥带走了
	// 我是口红，我给外卖小哥带走了
}
