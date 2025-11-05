package main

import "fmt"

/*

Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

*/

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{
		Root: &TrieNode{Children: make(map[rune]*TrieNode)},
	}
}

func (s *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := s.Root
	fmt.Println("insert")

	for _, ch := range word {
		if next, ok := node.Children[ch]; ok {
			fmt.Printf("found child for %s\n", string(ch))
			node = next
		} else {
			fmt.Printf("did not find child for %s\n", string(ch))
			newNode := &TrieNode{
				Children: make(map[rune]*TrieNode),
			}

			node.Children[ch] = newNode
			node = newNode
		}
	}
	node.IsEnd = true
}

func (s *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	node := s.Root
	fmt.Println("search")

	// a? -> p? -> p? -> l? -> e? isEnd?
	for _, ch := range word {
		if node == nil {
			return false
		}

		if next, ok := node.Children[ch]; ok {
			fmt.Printf("found child for %s\n", string(ch))
			node = next
		} else {
			fmt.Printf("did not find child for %s\n", string(ch))
			// does not exist ->
			return false
		}
	}

	return node.IsEnd
}

func (s *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	node := s.Root
	fmt.Println("startswith")

	// a? -> p? -> p? -> l? -> e? isEnd?
	for _, ch := range prefix {
		if node == nil {
			return false
		}

		if _, ok := node.Children[ch]; !ok {
			fmt.Printf("did not find child for %s\n", string(ch))
			return false
		}
		fmt.Printf("found child for %s\n", string(ch))
		node = node.Children[ch]

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
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
}
