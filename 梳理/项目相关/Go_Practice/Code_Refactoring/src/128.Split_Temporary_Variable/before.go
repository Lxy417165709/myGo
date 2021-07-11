package _28_Split_Temporary_Variable

import "math"

type Object1 struct {
	primaryForce   float64
	secondaryForce float64
	mass           float64
	delay          float64
}

func (o Object1) getDistanceTravelled(time float64) float64 {
	result := 0.0
	acc := o.primaryForce / o.mass // 第一次赋值
	primaryTime := math.Min(time, o.delay)
	result = 0.5 * acc * primaryTime
	secondaryTime := time - o.delay
	if secondaryTime > 0 {
		primaryVel := acc* o.delay
		acc = (o.primaryForce - o.secondaryForce) / o.mass // 第二次赋值
		result += primaryVel*secondaryTime + 0.5*acc*secondaryTime*secondaryTime
	}
	return result
}
