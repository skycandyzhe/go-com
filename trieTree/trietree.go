package trieTree

type trieNode struct {
	isWord bool // 是否是单词结尾
	next   map[rune]*trieNode
}

type trie struct {
	size int // 节点数量
	root *trieNode
}

func NewTrie() *trie {
	return &trie{
		0,
		&trieNode{false, make(map[rune]*trieNode)},
	}
}
func (u *trie) GetSize() int {
	return u.size
}

// 非递归算法
func (u *trie) Add(word string) {
	if len(word) == 0 {
		return
	}

	cur := u.root
	for _, v := range word {
		_, ok := cur.next[v] // 在NewTrie中已经初始化，能直接用
		if !ok {
			cur.next[v] = &trieNode{false, map[rune]*trieNode{}}
		}
		cur = cur.next[v]
	}
	// 判断该单词之前是否已经添加到tree中了
	if !cur.isWord {
		cur.isWord = true
		u.size++
	}
}

/*
 查询是否包含某个单词
 input  abc
 trieTree中存在 abcde
 返回 true
*/
func (u *trie) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}

	cur := u.root
	for _, v := range word {
		t1, ok := cur.next[v]
		if !ok {
			return false
		}
		cur = t1
	}
	return cur.isWord
}

/*
 前缀是否有以word为前缀的单词
 input  abc
 trieTree中存在 abcde
 返回 true
*/
func (u *trie) IsPrefix(word string) bool {
	if len(word) == 0 {
		return false
	}

	cur := u.root
	for _, v := range word {
		t1, ok := cur.next[v]
		if !ok {
			return false
		}
		cur = t1

	}
	return true
}

// 查看树中有没有对应的最长前缀
/*
 input  abcdee
 trieTree中存在 abc
 返回 true  abc
*/
func (u *trie) HasPrefix(word string) (bool, string) {
	if len(word) == 0 {
		// 空字符串没有前面
		return false, ""
	}
	flag := false
	var ret []rune
	laststr := ""
	cur := u.root
	for _, v := range word {
		t1, ok := cur.next[v]
		if !ok {
			return flag, laststr
			// return false, ""
		} else {
			ret = append(ret, v)
			if t1.isWord {
				flag = true
				laststr = string(ret)
			}
		}
		cur = t1

	}
	return flag, laststr
}
