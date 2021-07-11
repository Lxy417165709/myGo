package main

import (
	"fmt"
	"sort"
	"sync"
)

// ------------------------ MyMap ------------------------
type MyMap struct{
	Cap int
	Nodes []*Node	// 优化: 可以记录表头表尾
	mutex sync.Mutex
	// KeyDummyHead *KeyNode
	Keys []int
	hasBeenChannge bool
}

func NewMap(cap int) *MyMap{
	return &MyMap {
		Cap:cap,
		Nodes : make([]*Node,cap),
	}
}

func (m *MyMap) RangeAll() []int{
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.hasBeenChannge = false

	if !m.hasBeenChannge{
		return m.Keys
	}
	keys := make([]int,0)
	for _, nodeHead := range m.Nodes {
		curHead := nodeHead
		for curHead != nil{
			keys = append(keys,curHead.Key)
			curHead = curHead.Next
		}
	}
	sort.Ints(keys)
	m.Keys = keys
	return keys
}


func (m *MyMap) Insert(key,value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	insertIndex := key%m.Cap
	head := m.Nodes[insertIndex]
	insertNode :=  &Node{
		Key:key,
		Value:value,
	}
	if head == nil{
		m.Nodes[insertIndex] = insertNode
		return
	}
	for head.Next!=nil{
		// 更新操作
		if head.Key == key{
			if head.Value != value{
				head.Value = value
				m.hasBeenChannge = true
			}
			return
		}
		head = head.Next
	}
	head.Next = insertNode
	m.hasBeenChannge = true
}


// ------------------------ 底层结构 ------------------------
type Node struct{
	Key int
	Value interface{}
	Next *Node
}

type KeyNode struct{
	Key int
	Next *KeyNode
}

func main() {
	// 1-9,11-22...
	myMap := NewMap(10)
	for i := 1;i<10;i++{
		myMap.Insert(i,i*10+1)
	}
	fmt.Println(myMap.RangeAll())
}
