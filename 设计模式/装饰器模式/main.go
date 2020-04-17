package main

import (
	"fmt"
	"pattern"
)

func main() {

	// 第一种饮料
	var drink pattern.Drink = pattern.NaiCha{5}
	fmt.Printf("奶茶: %d\n",drink.CountPrice())

	drink = pattern.Pearl{Drink:drink,Price:2}
	fmt.Printf("奶茶 && 珍珠: %d\n",drink.CountPrice())

	drink = pattern.Pearl{Drink:drink,Price:2}
	fmt.Printf("奶茶 && 2 * 珍珠: %d\n",drink.CountPrice())

	drink = pattern.RedBean{Drink:drink,Price:10}
	fmt.Printf("奶茶 && 2 * 珍珠 && 红豆: %d\n",drink.CountPrice())

	// 第二种饮料

	var drink2 pattern.Drink = pattern.RedBean{pattern.RedBean{pattern.RedBean{pattern.PaoBing{10},2},2},2}

	fmt.Printf("刨冰 && 3 * 红豆: %d\n",drink2.CountPrice())

	// 奶茶: 5
	// 奶茶 && 珍珠: 7
	// 奶茶 && 2 * 珍珠: 9
	// 奶茶 && 2 * 珍珠 && 红豆: 19
	//
	// 刨冰 && 3 * 红豆: 16
}
