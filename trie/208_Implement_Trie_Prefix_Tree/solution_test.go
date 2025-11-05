package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")
	require.True(t, trie.Search("apple"))

	require.False(t, trie.Search("app"))
	require.True(t, trie.StartsWith("app"))

	trie.Insert("app")
	require.True(t, trie.Search("app"))

	// i also wanted to test a case when we have just app in trie. Should StartsWith(app) return true or false
}
