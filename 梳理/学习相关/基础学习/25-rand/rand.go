package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int31n(100))
	fmt.Println(rand.Int63n(100))
	rand.Float32()			// 取随机数[0, 1)，类型为float32
	rand.Float64()			// 取随机数[0, 1)，类型为float64
	// 23
	// 31
	// 80

}
