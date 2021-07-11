package pattern

type KissmeFactory struct {
}

func (f KissmeFactory) Create(name string) Food {
	switch name {
	case "naicha":
		return NaiCha{}
	case "snack":
		return Snack{}
	default:
		return nil
	}
}
