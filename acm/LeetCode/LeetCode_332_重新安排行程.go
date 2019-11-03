package main

import "fmt"

func findItinerary(tickets [][]string) []string {
	return solve(tickets,[]string{"JFK"})
}
func solve(tickets [][]string,ans []string) []string{

	now:=ans[len(ans)-1]

	m:=make(map[int]int)
	for i:=0;i< len(tickets);i++  {

	}
	return
}

func main() {
	fmt.Println(findItinerary([][]string{
		{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"},
	}))
}
