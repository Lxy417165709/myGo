package pattern


type Sad struct {
}

func (s Sad) ToString() string {
	return "我现在很伤心！"
}

func (s Sad)  GoToSendTakeOut() {
	me.state = Sad{}
}
func (s Sad) DrinkCaoMeiYiJunDuo() {
	me.state  = Happy{}
}

func (s Sad) BeComplained() {
	me.state = Sad{}
}
