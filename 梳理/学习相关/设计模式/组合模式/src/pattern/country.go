package pattern

import "fmt"

type Country struct {
	Name      string
	RegionMap map[string]Region
}

func NewCountry(name string) Country{
	return Country{name,make(map[string]Region)}
}

func (this Country) Insert(region Region) {
	this.RegionMap[region.GetName()] = region
}

func (this Country) Delete(name string){
	delete(this.RegionMap, name)
}
func (this Country) GetName() string {
	return this.Name
}
func (this Country) ShowRegionTree() {
	fmt.Printf("- %s - \n", this.GetName())
	for _, v := range this.RegionMap {
		v.ShowRegionTree()
	}
}
