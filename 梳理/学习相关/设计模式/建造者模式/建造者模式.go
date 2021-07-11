package main

import "pattern"

func main() {

	// 构建百香芒芒鲜莓
	// 新建构建指挥官
	buildDirector := pattern.BuildDirector{}

	// 设置构建者
	buildDirector.SetBuilder(pattern.BXMMXMBuilder{})

	// 指挥构建，获得实体
	bxmmxm := buildDirector.Build()

	// 展示信息
	bxmmxm.ShowInformation()




	// 构建波霸奶茶
	// 设置构建者
	buildDirector.SetBuilder(pattern.BobaNaiChaBuilder{})

	// 指挥构建，获得实体
	bobaNaiCha := buildDirector.Build()

	// 展示信息
	bobaNaiCha.ShowInformation()

	// 我是百香芒芒鲜莓，我采用 双享杯 杯。我加冰:100 g,加芒果: 50 g,加草莓: 50 g。我的盖子是: 双享盖
	// 我是波霸奶茶，我采用 大杯 杯。我加冰:250 g,加奶茶: 600 g。我的盖子是: 爱心盖
}

// 以上代码没有违背开闭原则，当有新的drink加入时，只需要加入一个相应的drinkBuilder类和一个实现饮料接口的实体类就可以了就可以了
