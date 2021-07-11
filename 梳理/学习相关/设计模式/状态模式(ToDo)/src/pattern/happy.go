package pattern

type Happy struct {
}

func (h Happy) ToString() string {
	return "我现在很开心！"
}

func (h Happy) GoToSendTakeOut() {
	me.state  = UnHappy{}
}
func (h Happy) DrinkCaoMeiYiJunDuo() {
	me.state  = Happy{}
}

func (h Happy) BeComplained() {
	me.state  = Sad{}
}
