package _25_valid_palindrome

func maxArea(height []int) int {
	res := 0

	// l = 0, r = last
	// look what is smaller and move that

	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] { //if left column is less than right
			l++ // shift the left pointer to the right
		} else {
			r-- // shift the right pointer to the left
		}
	}

	return res
}
