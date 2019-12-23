package main

import (
	"fmt"
	"sort"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	base:=folder[0]
	ans:=make([]string,0)
	ans = append(ans, base)
	for i:=1;i< len(folder);i++  {
		if !judge(base,folder[i]) {
			ans = append(ans, folder[i])
			base = folder[i]
		}
	}
	return ans
}
func judge(a,b string) bool{
	if len(a)> len(b) {
		a,b =b,a
	}
	for i:=0;i< len(a);i++  {
		if a[i]!=b[i] {
			return false
		}
	}
	if b[len(a)]=='/' {
		return true
	}

	return false
}

func main() {
	fmt.Println(removeSubfolders([]string{
		"/a","/a/b/c","/a/b/d","/b","/bb","/bbc","/b/cc","/bbc/cd/e",
	}))
}
