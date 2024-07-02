package main

import "testing"

func Test_isValid(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := isValid("()")
		if output != true {
			t.Fatalf("expected: %v, actual: %v", true, output)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		output := isValid("()[]{}")
		if output != true {
			t.Fatalf("expected: %v, actual: %v", true, output)
		}
	})
	t.Run("case 3", func(t *testing.T) {
		output := isValid("(]")
		if output != false {
			t.Fatalf("expected: %v, actual: %v", false, output)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		output := isValid("([)]")
		if output != false {
			t.Fatalf("expected: %v, actual: %v", false, output)
		}
	})

	t.Run("case 5", func(t *testing.T) {
		output := isValid("{[]}")
		if output != true {
			t.Fatalf("expected: %v, actual: %v", true, output)
		}
	})

	t.Run("case 6", func(t *testing.T) {
		output := isValid("([)]")
		if output != false {
			t.Fatalf("expected: %v, actual: %v", false, output)
		}
	})

	t.Run("case 7", func(t *testing.T) {
		output := isValid("([}}])")
		if output != false {
			t.Fatalf("expected: %v, actual: %v", false, output)
		}
	})
}
