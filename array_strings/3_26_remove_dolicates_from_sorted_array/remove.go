package main

// For example [1,1,2]
// returns the number of unique elements: 2, nums = [1,2,_]
func removeDuplicatesV1(nums []int) int {
	elements := map[int]bool{} //  O(n) space
	if len(nums) > 0 {
		for i := 0; i < len(nums); i++ { //  O(n)
			if _, ok := elements[nums[i]]; !ok {
				//  element is not in the map yet. Write it to the nums at respective index
				if i != 0 {
					index := len(elements)
					nums[index] = nums[i]
				}
				elements[nums[i]] = true
			}
		}
	}
	return len(elements)
}

// For example [1,1,2]
// returns the number of unique elements: 2, nums = [1,2,_]
func removeDuplicatesV2(nums []int) int {
	l := 0
	if len(nums) > 0 {
		for r := 1; r < len(nums); r++ {
			if nums[l] != nums[r] {
				l++
				nums[l] = nums[r]
			}
		}
	}
	return l + 1
}

// we can also assume that the duplicate is always next to the initial element
//  the first pair : 1,1. 1=1  so we skip i=0 j=0 		[1,1,2]
//  the second pair: 1,2. 1!=2 so we need to write the 2 to index 1 i=1 j=1	[1,2,2]

// we can also make the iterator look at the second element in the pair
