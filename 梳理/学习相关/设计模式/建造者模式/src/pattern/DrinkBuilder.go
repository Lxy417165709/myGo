package pattern



type DrinkBuilder interface {
	Prepare() DrinkBuilder
	Make() DrinkBuilder
	Cover() DrinkBuilder
	GetResult() Drink
}
