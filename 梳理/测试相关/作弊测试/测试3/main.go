package main

import "fmt"

type Person struct {
	Gang  bool // 0: 自己人 1: 敌人
	Power int  // 攻击力
	Hp    int
}

// p1 打 p2
// 打死则返回 true
func Hit(p1 *Person, p2 *Person) bool {
	p2.Hp -= p1.Power
	return p2.Hp <= 0
}

// 改变阵营
func CG(p1 *Person) {
	p1.Gang = !p1.Gang
}



func main() {
	p1 := &Person{false, 5, 100}
	p2 := &Person{true, 10, 100}
	fmt.Printf("你的基址是: %p\n",p1)
	fmt.Printf("他的基址是: %p\n",p2)
	hitMan:=0
	for {
		x := 0
		fmt.Scanf("%d\n", &x)
		switch x{
		case 1:
			switch Hit(p1,p2) {
			case true:
				hitMan++
				fmt.Printf("你已经杀死了 %d 个敌人\n", hitMan)
				if hitMan == 3 {
					fmt.Printf("你赢了\n")
					return
				}
				p2 = &Person{true, 10, 100} // 新敌人
			case false:
				fmt.Printf("你打了他 %d 血，他还有 %d 血\n", p1.Power,p2.Hp)
			}
		case 2:
			switch Hit(p2, p1) {
			case true:
				fmt.Printf("你输了\n")
				return
			case false:
				fmt.Printf("他打了你 %d 血，你还有 %d 血\n", p2.Power,p1.Hp)
			}
		case 3:
			CG(p1)
			fmt.Printf("你的阵营切换了\n")

		case 4:
			p1 = &Person{false, 5, 100}
			fmt.Printf("你的角色更新了\n")
			fmt.Printf("你的基址是: %p\n",p1)
		}

	}
}
