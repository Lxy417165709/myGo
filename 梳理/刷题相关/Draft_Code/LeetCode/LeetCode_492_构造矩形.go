package main

import "math"

func constructRectangle(area int) []int {
	W := int(math.Sqrt(float64(area)))
	for area%W != 0 {
		W--
	}
	return []int{area / W, W}
}
/*
	总结
	1. 这题太水..
*/
