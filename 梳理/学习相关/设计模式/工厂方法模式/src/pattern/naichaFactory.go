package pattern

type NaiChaFactory struct {
}

func (ncf NaiChaFactory) CreateFood() Food{
	return NaiCha{}
}
