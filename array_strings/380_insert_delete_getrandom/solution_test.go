package main

import (
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		s := Constructor()

		if !s.Insert(1) {
			t.Error("Expected true but got false")
		} // [1]

		if s.Remove(2) {
			t.Errorf("Remove(2) should return false")
		} // [1]

		if !s.Insert(2) {
			t.Errorf("Insert(2) should return true")
		} // [1,2]

		if val := s.GetRandom(); val != 2 && val != 1 {
			t.Errorf("GetRandom() should return 2 or 1, but got %d", val)
		}

		if !s.Remove(1) {
			t.Error("Expected true but got false")
		} // [2]

		if s.Insert(2) {
			t.Errorf("Insert(2) should return false")
		} // [2]

		if val := s.GetRandom(); val != 2 {
			t.Errorf("GetRandom() should return 2, but got %d", val)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		s := Constructor()

		if !s.Insert(0) {
			t.Fatalf("Expected true but got false")
		} // [0]
		t.Log(s.array)

		if !s.Insert(1) {
			t.Fatalf("Expected true")
		} // [0,1]
		t.Log(s.array, s.hashmap)

		if !s.Remove(0) {
			t.Fatalf("Expected true")
		} // [1]
		t.Log(s.array, s.hashmap)

		if !s.Insert(2) {
			t.Fatalf("Expected true")
		} // [1, 2]
		t.Log(s.array)
		t.Log(s.hashmap)

		if !s.Remove(1) { // why?
			t.Fatalf("Expected true")
		} // [2], actual [1]
		if s.array[0] != 2 {
			t.Fatalf("unexpected arr %v", s.array)
		}

		if val := s.GetRandom(); val != 2 {
			t.Errorf("GetRandom() should return 2, but got %d", val)
			return
		}
	})

	t.Run("case 3", func(t *testing.T) {
		s := Constructor()
		if ok := s.Remove(0); ok {
			t.Fatal("Expected false but got true")
		}
		if ok := s.Remove(0); ok {
			t.Fatal("Expected false but got true")
		}
		if ok := s.Insert(0); !ok {
			t.Fatal("Expected true but got false")
		}

		if val := s.GetRandom(); val != 0 {
			t.Fatalf("GetRandom() should return 0, but got %d", val)
		}

		t.Log(s.array, s.hashmap)

		if ok := s.Remove(0); !ok {
			t.Fatal("Expected true but got false")
		}

		t.Log(s.array, s.hashmap)

		if ok := s.Insert(0); !ok {
			t.Fatal("Expected true but got false")
		}
	})
}
