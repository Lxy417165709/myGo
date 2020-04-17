package pattern

type SnackFactory struct {
}

func (ncf SnackFactory) CreateFood() Food {
	return Snack{}
}
