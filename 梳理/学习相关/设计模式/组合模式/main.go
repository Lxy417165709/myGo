package main

import "pattern"

func main() {
	country := pattern.NewCountry("中国")
	province1 := pattern.NewProvince("广东省")
	city1 := pattern.NewCity("汕尾市")
	city2 := pattern.NewCity("河源市")
	province2 := pattern.NewProvince("福建省")
	city3 := pattern.NewCity("A市")
	city4 := pattern.NewCity("B市")
	city5 := pattern.NewCity("C市")
	province3 := pattern.NewProvince("浙江省")
	city6 := pattern.NewCity("杭州市")
	city7 := pattern.NewCity("黄山市")
	city8 := pattern.NewCity("G市")

	province1.Insert(city1)
	province1.Insert(city2)
	province2.Insert(city3)
	province2.Insert(city4)
	province2.Insert(city5)
	province3.Insert(city6)
	province3.Insert(city7)
	province3.Insert(city8)
	country.Insert(province1)
	country.Insert(province2)
	country.Insert(province3)

	country.ShowRegionTree()

	// - 中国 -
	//		广东省
	//			汕尾市
	//			河源市
	//		福建省
	//			A市
	//			B市
	//			C市
	//		浙江省
	//			杭州市
	//			黄山市
	//			G市


	country.Delete("浙江省")
	country.ShowRegionTree()
	// - 中国 -
	//		广东省
	//			河源市
	//			汕尾市
	//		福建省
	//			A市
	//			B市
	//			C市
}

// 以上就是组合模式了
// 组合模式很适用于树形结构
