package pattern

import "fmt"

type Drink interface {
	ShowInformation()
}

type BXMMXM struct {
	Bottle     string // 杯子
	Ice        int    // 冰量
	Mango      int    // 芒果量
	Strawberry int    // 草莓量
	Lid        string // 盖子
}

func (bx BXMMXM) ShowInformation() {
	fmt.Printf("我是百香芒芒鲜莓，我采用 %s 杯。我加冰:%d g,加芒果: %d g,加草莓: %d g。我的盖子是: %s\n",
		bx.Bottle, bx.Ice, bx.Mango, bx.Strawberry, bx.Lid,
	)
}

type BobaNaiCha struct {
	Bottle string // 杯子
	Ice    int    // 冰量
	NaiCha int    // 奶茶量
	Lid    string // 盖子
}

func (bbnc BobaNaiCha) ShowInformation() {
	fmt.Printf("我是波霸奶茶，我采用 %s 杯。我加冰:%d g,加奶茶: %d g。我的盖子是: %s\n",
		bbnc.Bottle, bbnc.Ice, bbnc.NaiCha, bbnc.Lid,
	)
}
