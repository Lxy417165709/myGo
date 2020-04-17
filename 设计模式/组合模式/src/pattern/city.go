package pattern

import "fmt"

// 当前的最小单位(这里不判断插入异常，这样可以方便拓展)
type City struct {
	Name      string
	RegionMap map[string]Region
}

func NewCity(name string) City {
	return City{name, make(map[string]Region)}
}

func (this City) Insert(region Region){
	this.RegionMap[region.GetName()] = region
}

func (this City) Delete(name string){
	delete(this.RegionMap, name)
}
func (this City) GetName() string {
	return this.Name
}
func (this City) ShowRegionTree() {
	fmt.Printf("		%s\n", this.GetName())
	for _, v := range this.RegionMap {
		v.ShowRegionTree()
	}
}
