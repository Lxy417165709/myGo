package pattern

type myself struct {
	state State
}
func (s *myself) ToString() string {
	return s.state.ToString()
}

func (s *myself) GoToSendTakeOut() {
	s.state.GoToSendTakeOut()
}
func (s *myself) DrinkCaoMeiYiJunDuo() {
	s.state.DrinkCaoMeiYiJunDuo()
}

func (s *myself) BeComplained() {
	s.state.BeComplained()
}

// 单例模式
// 初始化为开心状态
var me = &myself{Happy{}}

func GetMyself() *myself{
	return me
}
