package main

type WordDictionary struct {
	IsEnd bool
	Node  [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{false, [26]*WordDictionary{}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		pos := uint8(word[i] - 'a')
		if root.Node[pos] == nil {
			root.Node[pos] = &WordDictionary{false, [26]*WordDictionary{}}
		}
		root = root.Node[pos]
	}
	root.IsEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return solve(this, word)
}

func solve(root *WordDictionary, word string) bool {
	if root == nil {
		return false
	}
	if len(word) == 0 {
		return root.IsEnd
	}

	if word[0] == '.' {
		for i := 0; i < len(root.Node); i++ {
			if solve(root.Node[i], word[1:]) {
				return true
			}
		}
		return false
	} else {
		pos := word[0] - 'a'
		return solve(root.Node[pos], word[1:])
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

/*
	总结
	1. 这题和208差不多，不过查找单词时加了一个 '.'符号。
*/
