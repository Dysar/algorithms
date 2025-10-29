package main

/*
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.


Input: ratings = [1,0,2]
Output: 5

Input: ratings = [1,2,2]
Output: 4
*/

func candy(ratings []int) int {

	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}

	candies := make([]int, len(ratings))
	candies[0] = 1 // init
	for i := 1; i < len(ratings); i++ {
		candies[i] = 1 // init

		// is the rating at i greater than the rating at i-1? -> then it needs more candy
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// go from pre-last value to the first
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	total := 0
	for _, candyNum := range candies {
		total += candyNum
	}

	return total
}
