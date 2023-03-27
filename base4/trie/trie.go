package trie

type Trie struct {
	nest  [26]*Trie
	isEnd bool
}


func Constructor() Trie {
	return Trie{}
}


func (this *Trie) Insert(word string)  {
	bytes := []byte(word)

	for _, v := range bytes {
		i := v - 'a'
		if this.nest[i] == nil {
			this.nest[i] = &Trie{}
		}
		this = this.nest[i]
	}
	this.isEnd = true
}


func (this *Trie) Search(word string) bool {
	bytes := []byte(word)

	for _, v := range bytes {
		i := v - 'a'
		if this.nest[i] == nil {
			return false
		}
		this = this.nest[i]
	}
	if this.isEnd {
		return true
	}
	return false
}


func (this *Trie) StartsWith(prefix string) bool {
	bytes := []byte(prefix)

	for _, v := range bytes {
		i := v - 'a'
		if this.nest[i] == nil {
			return false
		}
		this = this.nest[i]
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