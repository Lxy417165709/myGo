package main

import "strings"

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Index(A+A, B) != -1
}

/*
	总结
	1. 这个解法太牛逼了...
	2. 除了牛逼我无话可说！
*/
