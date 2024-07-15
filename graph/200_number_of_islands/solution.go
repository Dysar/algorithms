package main

// Golang queue: need + buffer.
// c := make(chan int, 300)
// Push
// c <- value
// Pop
// x <- c

func numIslands(grid [][]byte) int {
	// input validation
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// get the dimensions of the grid
	rows, cols := len(grid), len(grid[0])

	// we want to mark cells visited
	visited := make([][]bool, rows) // SC O(n)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// we want to count the number of islands
	islands := 0

	bfs := func(r, c int) {
		q := make(chan [2]int, rows*cols)
		visited[r][c] = true
		q <- [2]int{r, c}

		for len(q) > 0 {
			queued := <-q
			directions := [][]int{
				// {horizontal, vertical}
				{1, 0},  // direction to the right
				{-1, 0}, // direction to the left
				{0, 1},  // direction above
				{0, -1}, // direction below
			}

			for _, d := range directions {

				dr, dc := d[0], d[1]
				r, c := queued[0]+dr, queued[1]+dc
				// check that this position is in bounds
				if r < rows && r >= 0 && c < cols && c >= 0 &&
					// this position is land position
					grid[r][c] == '1' &&
					// this position hasn't been visited
					!visited[r][c] {

					// add this position to the queue to run BFS on this cell
					q <- [2]int{r, c}
					// and mark the position visited
					visited[r][c] = true
				}
			}
		}
	}

	// let's visit every position in the grid
	for r := 0; r < rows; r++ { // TC O(n):
		for c := 0; c < cols; c++ {
			// if we visit '0' we don't have to do anything, but if we visit a 1,
			// then we have to traverse it and mark it visited
			if grid[r][c] == '1' && !visited[r][c] {
				bfs(r, c)
				islands++ // only increment if we get to the one we haven't already visited
			}
		}
	}

	return islands
}
