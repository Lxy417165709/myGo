package _28_Split_Temporary_Variable

import "math"

type Object2 struct {
	primaryForce   float64
	secondaryForce float64
	mass           float64
	delay          float64
}

// 重构1
func (o Object2) getDistanceTravelled(time float64) float64 {
	result := 0.0
	primaryAcc := o.primaryForce / o.mass
	primaryTime := math.Min(time, o.delay)
	result = 0.5 * primaryAcc * primaryTime
	secondaryTime := time - o.delay
	if secondaryTime > 0 {
		primaryVel := primaryAcc * o.delay
		secondAcc := (o.primaryForce - o.secondaryForce) / o.mass // 第二次赋值
		result += primaryVel*secondaryTime + 0.5*secondAcc*secondaryTime*secondaryTime
	}
	return result
}

// 重构2
func (o Object2) getDistanceTravelled2(time float64) (result float64) {
	result += 0.5 * o.getPrimaryAcc() * o.getPrimaryTime(time)
	if o.getSecondTime(time) > 0 {
		result += (o.getPrimaryVel() + 0.5*o.getSecondAcc()*o.getSecondTime(time)) * o.getSecondTime(time)
	}
	return
}

func (o Object2) getSecondTime(time float64) float64 {
	return time - o.delay
}
func (o Object2) getPrimaryTime(time float64) float64 {
	return math.Min(time, o.delay)
}
func (o Object2) getPrimaryVel() float64 {
	return o.getAcc(o.primaryForce, 0, o.mass) * o.delay
}
func (o Object2) getPrimaryAcc() float64 {
	return o.getAcc(o.primaryForce, 0, o.mass)
}
func (o Object2) getSecondAcc() float64 {
	return o.getAcc(o.primaryForce, o.secondaryForce, o.mass)
}
func (o Object2) getAcc(primaryForce, secondaryForce, mass float64) float64 {
	return (primaryForce - secondaryForce) / mass
}
