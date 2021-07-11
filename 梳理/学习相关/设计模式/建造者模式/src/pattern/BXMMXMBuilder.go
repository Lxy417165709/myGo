package pattern

// 百香芒芒鲜莓
type BXMMXMBuilder struct {
	bxmmxm BXMMXM
}

func (bxb BXMMXMBuilder) Prepare() DrinkBuilder {
	bxb.bxmmxm.Bottle = "双享杯"
	return bxb
}
func (bxb BXMMXMBuilder) Make() DrinkBuilder {
	bxb.bxmmxm.Mango = 50
	bxb.bxmmxm.Strawberry = 50
	bxb.bxmmxm.Ice = 100
	return bxb
}
func (bxb BXMMXMBuilder) Cover() DrinkBuilder {
	bxb.bxmmxm.Lid = "双享盖"
	return bxb
}

func (bx BXMMXMBuilder) GetResult() Drink {
	return bx.bxmmxm
}
