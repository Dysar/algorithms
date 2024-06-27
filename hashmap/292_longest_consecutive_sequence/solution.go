package main

// it must be  O(n)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := map[int]bool{}
	longestSequence := 1
	for _, num := range nums {
		set[num] = true
	}

	for _, num := range nums {
		// check if the element has next element. Then it has potential to be the start of the sequence
		if _, ok := set[num+1]; !ok {
			// if no, continue
			continue
		}

		// if yes, check if the element has the previous element.
		if _, ok := set[num-1]; ok {
			// if yes, it's not the start of the sequence, continue
			continue
		}
		// if no -> it's a start of the sequence.
		// We already know that it has a next element, so let's check after that
		sequenceLen := 2
		nextNum := num + 2
		for {
			if _, ok := set[nextNum]; !ok {
				break
			} else {
				sequenceLen++
			}
			nextNum++
		}

		longestSequence = max(longestSequence, sequenceLen)
	}

	// we visit each element at most twice. therefore, it's TC O(n)
	// we have a map (hashset) -> SC O (n)
	return longestSequence
}
