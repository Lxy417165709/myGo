package pattern


type UnHappy struct {
	nowState State
}

func (uh  UnHappy) ToString() string {
	return "我现在很不开心！"
}

func (uh  UnHappy) GoToSendTakeOut() {
	me.state = Sad{}
}
func (uh  UnHappy) DrinkCaoMeiYiJunDuo() {
	me.state = Happy{}
}

func (uh  UnHappy) BeComplained() {
	me.state = Sad{}
}
