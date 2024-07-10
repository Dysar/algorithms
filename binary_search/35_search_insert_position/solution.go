package _5_search_insert_position

// Given a sorted array of distinct integers and a target value,
// return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
func searchInsert(nums []int, target int) int {
	// cut the slice in half, compare to the value
	l, r := 0, len(nums)-1 // 0,3
	for l <= r {

		// calculate the middle
		m := (l + r) / 2 //0+3/2 = 1
		val := nums[m]   //[1,3,5,6] -> 3
		if target > val {
			l = m + 1 // 2,3
		} else if target == val {
			return m
		} else if target < val {
			r = m - 1 // 0,0
		}
	}

	return l
}
