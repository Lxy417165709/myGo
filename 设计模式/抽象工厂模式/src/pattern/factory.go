package pattern

type Factory interface{
	CreateFood() Food
	CreateCar() Car
}
