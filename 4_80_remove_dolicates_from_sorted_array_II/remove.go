package main

// For example nums = [1,1,1,2,2,3]
// Output 5, nums = [1,1,2,2,3,_]
func removeDuplicatesV1(nums []int) int {
	// l is left, m is middle, r is right
	l := 0
	m := 1
	r := 2

	if len(nums) > 0 {
		for r < len(nums) {
			// if we have 1,1,1, we have to do something
			if nums[l] == nums[r] {
				// this should look at the array starting from index 3
				// [1,1,1,{2}...
				for i := r + 1; i < len(nums); i++ {
					if nums[i] != nums[m] {

					}
				}
				// if we are here we did not find any different elements
				return 2 // k will be 2 because our result is [1,1]
			}
		}
	}
	return l + 1
}

// here instead of comparing pairs we can compare 3 elements
// what are the options: 1,1,1; 1,2,2; 1,2,3.
// if we have 1,1,1, we have to do something.
//	For example look further for elements that have different values.
//  if we don't find such an element, return immediately
//	if we found such an element, put that to the 3rd place and
// 	set the cursor to the 3rd place for the next iteration
// if we have 1,2,2 or 1,2,3 it's essentially ok

// For example nums = [1,1,1,2,2,3]
// Output 5, nums = [1,1,2,2,3,_]
func removeDuplicatesV2(nums []int) int {
	// slow index, fast index, counter
	slow, fast := 0, 0

	// for 0 or 1 or 2 len
	if len(nums) <= 2 {
		return len(nums)
	}

	// len is at least 2
	for fast < len(nums) { // if len nums 6, last index is 5. 5<6 => ok
		count := 1
		for fast+1 < len(nums) && nums[fast] == nums[fast+1] {
			fast += 1
			count += 1
		}
		for i := 1; i <= min(2, count); i++ {
			nums[slow] = nums[fast]
			slow += 1
		}
		fast += 1
	}

	return slow
}

// two pointer approarch
// we set the 2 pointers:
// result: [{{1}},1,1,2,2,3], fast = 0, slow = 0, count = 1

// iter 1st
// 1+1< 6 and 1=1 => fast = 1; count = 2;  2+1<6 and 1=1 => fast = 2; count = 3

//nums[0] = nums[2], slow = 1; nums[1] = nums[2], slow = 2
