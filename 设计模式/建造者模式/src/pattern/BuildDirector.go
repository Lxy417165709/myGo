package pattern

type BuildDirector struct {
	Builder DrinkBuilder
}

// 构建过程透明化
func (bd BuildDirector) Build() Drink {
	return bd.Builder.Prepare().Make().Cover().GetResult()
}

func (bd *BuildDirector) SetBuilder(drinkBuilder DrinkBuilder) {
	bd.Builder = drinkBuilder
}
