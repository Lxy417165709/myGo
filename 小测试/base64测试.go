package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	bytes,_ := ioutil.ReadFile("acm/LeetCode/LeetCode_1_两数之和.go")

	base64Str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(base64Str)
	// ok!
}
