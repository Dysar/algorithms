package _25_valid_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		is := isPalindrome("A man, a plan, a canal: Panama")
		if !is {
			t.Fatalf("isPalindrome returned false")
		}
	})
	t.Run("case 2", func(t *testing.T) {
		is := isPalindrome("race a car")
		if is {
			t.Fatalf("isPalindrome returned true")
		}
	})
	t.Run("case 3", func(t *testing.T) {
		is := isPalindrome(" ")
		if !is {
			t.Fatalf("isPalindrome returned false")
		}
	})

	t.Run("case 4", func(t *testing.T) {
		is := isPalindrome("0P")
		if is {
			t.Fatalf("isPalindrome returned true")
		}
	})

	t.Run("case 5", func(t *testing.T) {
		is := isPalindrome("a")
		if !is {
			t.Fatalf("expected true")
		}
	})
}
