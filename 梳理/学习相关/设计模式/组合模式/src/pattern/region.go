package pattern



type Region interface {
	Insert(region Region)
	Delete(name string)
	GetName() string
	ShowRegionTree()
}
