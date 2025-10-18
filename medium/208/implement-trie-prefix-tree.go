package leetcode

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:
- Trie() Initializes the trie object.
- void insert(String word) Inserts the string word into the trie.
- boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
- boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/
type Trie struct {
	childs [26]*Trie
	isEnd  bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	current := this

	for _, v := range word {
		id := v - 'a'
		if current.childs[id] == nil {
			current.childs[id] = &Trie{}
		}
		current = current.childs[id]
	}

	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this

	for _, v := range word {
		id := v - 'a'
		if current.childs[id] == nil {
			return false
		}
		current = current.childs[id]
	}

	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this

	for _, v := range prefix {
		id := v - 'a'
		if current.childs[id] == nil {
			return false
		}
		current = current.childs[id]
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
