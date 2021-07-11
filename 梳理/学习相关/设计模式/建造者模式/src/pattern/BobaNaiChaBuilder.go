package pattern

// 波霸奶茶创建器
type BobaNaiChaBuilder struct {
	bbnc BobaNaiCha
}

func (bbncb BobaNaiChaBuilder) Prepare() DrinkBuilder {
	bbncb.bbnc.Bottle = "大杯"
	return bbncb
}
func (bbncb BobaNaiChaBuilder) Make() DrinkBuilder {
	bbncb.bbnc.Ice = 250
	bbncb.bbnc.NaiCha = 600
	return bbncb
}
func (bbncb BobaNaiChaBuilder) Cover() DrinkBuilder {
	bbncb.bbnc.Lid = "爱心盖"
	return bbncb
}

func (bbncb BobaNaiChaBuilder) GetResult() Drink {
	return bbncb.bbnc
}
