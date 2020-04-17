package pattern

type KissmeFactory struct {
}

func (kmf KissmeFactory) CreateFood() Food{
	return NaiCha{}
}

func (kmf KissmeFactory) CreateCar() Car{
	return BaoMa{}
}
