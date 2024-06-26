package main

import "testing"

func Test_isHappy(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		isHappy := isHappy(91)
		if isHappy != true {
			t.Fatalf("isHappy = %v, want %v", isHappy, true)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		t.Log(isHappy(8282))
	})
	t.Run("case 3", func(t *testing.T) {
		isHappy := isHappy(2)
		if isHappy != false {
			t.Fatalf("must be false, but %v", isHappy)
		}
	})
	t.Run("case 4", func(t *testing.T) {
		isHappy := isHappy(1)
		if isHappy != true {
			t.Fatalf("must be true, but %v", isHappy)
		}
	})

}

func TestBasic(t *testing.T) {
	t.Log(19 / 10)
	t.Log(19 % 10)

	t.Log(1 / 10)

	t.Log(8282 / 10)
	t.Log(8282 % 10)

	t.Log(828 / 10)
	t.Log(828 % 10)
}
