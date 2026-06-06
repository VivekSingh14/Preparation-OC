package main

func main2() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}

	setZeroes(matrix)

}

// Wrong solution
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])

	rowZero := false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				if i != 0 {
					matrix[i][0] = 0
					continue
				}
				rowZero = true
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}

	if rowZero {
		for i := 0; i < cols; i++ {
			matrix[0][i] = 0
		}
	}

}

// right solution
func setZeroes1(matrix [][]int) {
	var (
		rows         = len(matrix)
		cols         = len(matrix[0])
		firstRowFlag = false
	)

	// loop through the matrix, and using elements of the first row and the first column as flags
	// first row -> determine column
	// first column -> determine row
	// the index[0][0] belongs to the first row -> extra variable to determine if the first row need to be set to 0's
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				if r != 0 {
					matrix[r][0] = 0
					continue
				}
				firstRowFlag = true
			}
		}
	}

	// loop through the matrix, except for the first row and the first column
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// the first column
	if matrix[0][0] == 0 {
		for r := 1; r < rows; r++ {
			matrix[r][0] = 0
		}
	}

	// the first row
	if firstRowFlag {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}

}
