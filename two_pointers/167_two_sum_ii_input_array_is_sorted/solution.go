package _25_valid_palindrome

func twoSum(numbers []int, target int) []int {

	// numbers is 1-indexed array of integers numbers that is already sorted in non-decreasing order

	// find 2 numbers that add up to target

	// return the indices of two numbers

	// there is exactly one solution. you cannot use same element twice
	// only constant space

	l, r := 0, 1
	// l can go up to pre-last element.  0,1,2; 3 -> l < 3-1.
	// r can go up to the last element.  0,1,2; 3 -> r < 3.
	for l < len(numbers)-1 && r < len(numbers) {
		// fmt.Printf("l:%d, r:%d\n", l, r)
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else {
			// should go 0,1 -> 0,2 -> 0,3, 1,2, 1,3 -> 2,3
			if r == len(numbers)-1 {
				l++
				r = l + 1
			} else {
				r++
			}
		}
	}
	return []int{l + 1, r + 1}
}
