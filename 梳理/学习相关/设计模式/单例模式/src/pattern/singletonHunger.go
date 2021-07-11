package pattern

// 饿汉式
type HungerSingleton struct {
}

var hungerInstance = & HungerSingleton{}
func GetHungerInstance() * HungerSingleton {
	return hungerInstance
}
