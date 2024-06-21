package main

import "testing"

func Test_groupAnagrams(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
		if len(result) != 3 {
			t.Fatalf("result should have length 3, got %d", len(result))
		}
		t.Log(result)
	})

	t.Run("case 2", func(t *testing.T) {
		result := groupAnagrams([]string{})
		if len(result) != 0 {
			t.Fatalf("result should be empty")
		}
	})

	t.Run("case 3", func(t *testing.T) {
		result := groupAnagrams([]string{"a"})
		if result[0][0] != "a" {
			t.Fatalf("expected: %v, got: %v", "a", result[0][0])
		}
	})
}
