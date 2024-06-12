package __169_majority_element

// majorityElement uses Moore Voting Algorithm
func majorityElement(nums []int) int {
	cnt, m := 0, 0

	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			m = nums[i]
			cnt = 1
		} else {
			if nums[i] == m {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return m
}
