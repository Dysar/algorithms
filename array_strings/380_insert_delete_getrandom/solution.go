package main

import "math/rand"

type RandomizedSet struct {
	// hashmap will map the values to indices
	hashmap map[int]int
	array   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashmap: make(map[int]int),
		array:   make([]int, 0)}
}

// Insert average O(1)
func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.hashmap[val]; ok {
		return false
	}
	// Complexity ??
	s.array = append(s.array, val)
	// we have a map [], we add 1. len(s.array) is now 1, index of element is 0
	s.hashmap[val] = len(s.array) - 1
	return true
}

// Remove average O(1)
func (s *RandomizedSet) Remove(val int) bool {
	if len(s.array) == 0 || len(s.hashmap) == 0 {
		return false
	}

	index, ok := s.hashmap[val]
	if !ok {
		return false
	}

	// copy last value to the index we are removing from
	lastElement := s.array[len(s.array)-1]
	s.array[index] = lastElement

	// and pop the last element
	s.array = s.array[:len(s.array)-1]

	// now update the index of the moved element in the map
	if _, ok := s.hashmap[lastElement]; ok {
		s.hashmap[lastElement] = index
	}

	delete(s.hashmap, val)

	return true
}

// GetRandom average O(1)
func (s *RandomizedSet) GetRandom() int {
	n := len(s.hashmap)
	// get a random number [0,n)
	randIndex := rand.Intn(n)
	return s.array[randIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
