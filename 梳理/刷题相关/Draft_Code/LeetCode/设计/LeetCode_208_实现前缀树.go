package main

type Trie struct {
	Char  byte // 这个字段可以删除的
	IsEnd bool
	Node  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{0, true, [26]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		pos := uint8(word[i] - 'a')
		if root.Node[pos] == nil {
			root.Node[pos] = &Trie{word[i], false, [26]*Trie{}}
		}
		root = root.Node[pos]
	}
	root.IsEnd = true;
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		pos := uint8(word[i] - 'a')
		if root.Node[pos] == nil {
			return false
		}
		root = root.Node[pos]
	}
	return root.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := 0; i < len(prefix); i++ {
		pos := uint8(prefix[i] - 'a')
		if root.Node[pos] == nil {
			return false
		}
		root = root.Node[pos]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
