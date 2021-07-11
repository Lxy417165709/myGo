package pattern

import "fmt"

type TV struct{}

func(tv TV)Open() {
	fmt.Println("电视机打开了")
}

func(tv TV)Close() {
	fmt.Println("电视机关闭了")
}
