package trieTree

import (
	"fmt"
	"testing"
)

func TestTRIE(t *testing.T) {
	root := NewTrie()
	root.Add("a")
	root.Add("b")
	root.Add("baaa")
	root.Add("apple")
	root.Add("apple watch")
	root.Add("banana")
	root.Add("water")
	root.Add("c")
	root.Add("mm")
	root.Add("你好世界")

	// root.traversal(0)

	testCases := [...]string{
		"ba", "water", "watermelon",
		"app", "apple tv",
		"mango", "你好世界", "你好朋友", "你好世界上的朋友",
	}

	for _, str := range testCases {

		has, pre := root.HasPrefix(str)
		fmt.Println(str+" isMatch ", has, pre)
	}
}
