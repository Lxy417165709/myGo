package pattern


// 当前状态
type NowState struct {
	state State
}

func (ns *NowState) ToString() string {
	return ns.state.ToString()
}

func (ns *NowState) GoToSendTakeOut() {
	ns.state.GoToSendTakeOut()
}
func (ns *NowState) DrinkCaoMeiYiJunDuo() {
	ns.state.DrinkCaoMeiYiJunDuo()
}

func (ns *NowState) BeComplained() {
	ns.state.BeComplained()
}
