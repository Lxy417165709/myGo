package pattern

type BakechaFactory struct {
}

func (bkf BakechaFactory) CreateFood() Food{
	return Snack{}
}

func (bkf BakechaFactory) CreateCar() Car{
	return Suv{}
}
