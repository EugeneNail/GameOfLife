package internal

func createGrid(rows int, columns int) [][]bool {
	var grid [][]bool

	for row := 0; row < rows; row++ {
		var cells []bool
		for column := 0; column < columns; column++ {
			cells = append(cells, false)
		}
		grid = append(grid, cells)
	}

	return grid
}

func boolToInt(value bool) int {
	// Looks dumb, but it's a way more performant than the "clever" variant
	// https://0x0f.me/blog/golang-compiler-optimization/
	var i int
	if value {
		i = 1
	} else {
		i = 0
	}
	return i
}
