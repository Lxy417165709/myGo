package main

type newType map[int]int
func (nt *newType) insert(a,b int) {
	(map[int]int)(*nt)[a] = b
}
func main()  {
	a := new(newType)
	a.insert(1,2)
	// panic: assignment to entry in nil map
}
