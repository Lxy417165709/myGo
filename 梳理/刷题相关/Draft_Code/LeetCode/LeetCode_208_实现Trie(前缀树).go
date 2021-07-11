package main

type Trie struct {
	char   uint8     // 这个实际没有什么作用 0.0
	isWord bool      // 0表示不是单词，1表示是单词
	son    [26]*Trie // 0,1,2.. 表示 'a','b','c'...
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{100, false, [26]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		if this.son[word[i]-'a'] == nil {
			this.son[word[i]-'a'] = &Trie{word[i] - 'a', false, [26]*Trie{}}
		}
		this = this.son[word[i]-'a']
	}
	this.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if this.son[word[i]-'a'] == nil {
			return false
		}
		this = this.son[word[i]-'a']
	}
	return this.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if this.son[prefix[i]-'a'] == nil {
			return false
		}
		this = this.son[prefix[i]-'a']
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
func main() {

}

/*
	总结
	1. 这是一道前缀树的题，第一次的时候提交AC了，但是代码还不算太优美
	   于是看了官方解答，把代码优美了一下 0.0
*/
