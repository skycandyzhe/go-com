package trieTree

import (
	"fmt"
	"testing"
)

func TestTRIE(t *testing.T) {
	root := NewTrie()
	root.Add("a")
	root.Add("b")
	root.Add("ba")
	root.Add("baaa")
	root.Add("apple")
	root.Add("apple watch")
	root.Add("banana")
	root.Add("wat")
	root.Add("water")
	root.Add("c")
	root.Add("mm")
	root.Add("你好")
	root.Add("你好世界")
	root.Add("你好世界")
	root.Add("你好世界1")
	root.Add("你好世界12")
	root.Add("你好世界123")
	root.Add("你好世界1234")

	// root.traversal(0)

	testCases := [...]string{
		"ba", "water", "watermelon", "你好世界", "你好世界1234", "你好世界上的朋友",
	}

	for _, str := range testCases {

		pre := root.HasPrefixs(str)
		fmt.Println(str+" isMatch ", pre)
	}
}
