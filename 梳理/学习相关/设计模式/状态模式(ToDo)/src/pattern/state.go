package pattern

type State interface {
	ToString() string     // 字符串化，方便查看
	GoToSendTakeOut()     // 去送外卖(操作)  --> 这个操作会影响状态
	DrinkCaoMeiYiJunDuo() // 喝草莓益菌多(操作) --> 这个操作会影响状态
	BeComplained()        // 被投诉(操作) --> 这个操作会影响状态
}







