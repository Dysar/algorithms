package main

// [1,2,3,4,5], k = 2; [4,5,1,2,3]
func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse := func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l, r = l+1, r-1
		}
	}
	// reverse the array to [5,4,3,2,1]
	l, r := 0, len(nums)-1
	reverse(l, r)

	// reverse the first k elements to [4,5,3,2,1]
	l, r = 0, k-1
	reverse(l, r)

	// reverse the remaining elements [4,5,1,2,3]
	l, r = k, len(nums)-1
	reverse(l, r)
}
