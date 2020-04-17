package pattern

import "fmt"

type Province struct {
	Name      string
	RegionMap map[string]Region
}

func NewProvince(name string) Province{
	return Province{name,make(map[string]Region)}
}

func (this Province) Insert(region Region){
	this.RegionMap[region.GetName()] = region
}

func (this Province) Delete(name string) {
	delete(this.RegionMap, name)
}
func (this Province) GetName() string {
	return this.Name
}
func (this Province) ShowRegionTree() {
	fmt.Printf("	%s\n", this.GetName())
	for _, v := range this.RegionMap {
		v.ShowRegionTree()
	}
}
