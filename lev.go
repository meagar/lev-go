package lev

// Distance computes and returns the Levenshtein Distance between two words
// See https://en.wikipedia.org/wiki/Levenshtein_distance
func Distance(a, b string) int {
	return singleRowDistance(a, b)
}

// DistanceD computes and returns the Damerau–Levenshtein distance between two words.
// See https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
// The Damerau-Levenshtein algorithm also allows for transposing two adjacent letters;
// where The Levenshtein distance from "ab" to "ba" would be 2 (one deletion, one addition),
// the Damerau-Levenshtein distance is 1, a single transposition/swap of the two letters.
func DistanceD(a, b string) int {
	return matrixDistanceD(a, b)
}

func naiveDistance(a, b string) int {
	// All implementations share the same pre-conditions:
	// If the length of either word is 0, then the length of the opposite word is the edit distance
	if len(b) == 0 {
		return len(a)
	}
	if len(a) == 0 {
		return len(b)
	}

	if a[0] == b[0] {
		return naiveDistance(a[1:], b[1:])
	}

	// The naive option recurses down three paths and takes the smallest value, and adds 1 to it
	return 1 + min3(
		naiveDistance(a[1:], b),
		naiveDistance(a, b[1:]),
		naiveDistance(a[1:], b[1:]),
	)
}

// matrixDistance uses a nxm matrix, where n and m are the length of the words.
// Each position in the matrix represents the Lev distance of the corresponding sub-strings.
// The top row and left column are "seeded" with 1..n and 1..m, as each of these cells represents
// the comparison against an empty sub-string.
func matrixDistance(a, b string) int {
	n, m := len(a), len(b)
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	matrix := make([][]int, m+1)
	for i := 0; i < len(b)+1; i++ {
		matrix[i] = make([]int, n+1)
	}

	for i, row := range matrix {
		row[0] = i
	}

	for i := range matrix[0] {
		matrix[0][i] = i
	}

	for y := 1; y < n+1; y++ {
		for x := 1; x < m+1; x++ {
			if b[x-1] == a[y-1] {
				matrix[x][y] = matrix[x-1][y-1]
			} else {
				matrix[x][y] = 1 + min3(
					matrix[x-1][y],
					matrix[x][y-1],
					matrix[x-1][y-1],
				)
			}
		}
	}

	return matrix[m][n]
}

// matrixDistance uses a nxm matrix, where n and m are the length of the words.
// Each position in the matrix represents the Lev distance of the corresponding sub-strings.
// The top row and left column are "seeded" with 1..n and 1..m, as each of these cells represents
// the comparison against an empty sub-string.
func matrixDistanceD(a, b string) int {
	n, m := len(a), len(b)
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	matrix := make([][]int, m+1)
	for i := 0; i < len(b)+1; i++ {
		matrix[i] = make([]int, n+1)
	}

	for i, row := range matrix {
		row[0] = i
	}

	for i := range matrix[0] {
		matrix[0][i] = i
	}

	for y := 1; y < n+1; y++ {
		for x := 1; x < m+1; x++ {
			if b[x-1] == a[y-1] {
				matrix[x][y] = matrix[x-1][y-1]
			} else {
				matrix[x][y] = 1 + min3(
					matrix[x-1][y],
					matrix[x][y-1],
					matrix[x-1][y-1],
				)

				if x > 1 && y > 1 && a[y-1] == b[x-2] && a[y-2] == b[x-1] {
					matrix[x][y] = min2(
						matrix[x][y],
						matrix[x-2][y-2]+1) // transposition
				}
			}
		}
	}

	return matrix[m][n]
}

// Double row is an optimization over the matrix implementation, where only two rows of the matrix are
// actually instantiated in memory. The full matrix is simulated by repeatedly swapping which row is
// being filled, and which row is considered the previous row.
func doubleRowDistance(a, b string) int {
	n, m := len(a), len(b)

	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	row1 := make([]int, n+1)
	row2 := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		row2[i] = i
	}

	for y := 0; y < m; y++ {
		row1, row2 = row2, row1
		row2[0] = y + 1
		for x := 0; x < n; x++ {
			if a[x] == b[y] {
				row2[x+1] = row1[x]
			} else {
				row2[x+1] = 1 + min3(
					row1[x+1],
					row1[x],
					row2[x],
				)
			}
		}
	}

	return row2[len(row2)-1]
}

// The single row version is a further optimization where only one row is instantiated, and the
// new values are written over the row in-place. Only one temporary variable is required to maintain
// enough of the previous state to simulate a two-row implementation (which in turn effectively
// simulates the matrix implementation)
func singleRowDistance(a, b string) int {
	n, m := len(a), len(b)

	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	row := make([]int, n+1)

	for i := range row {
		row[i] = i
	}

	var last int

	for y := 0; y < m; y++ {
		last, row[0] = row[0], y+1
		for x := 0; x < n; x++ {
			if a[x] == b[y] {
				last, row[x+1] = row[x+1], last
			} else {
				last, row[x+1] = row[x+1], 1+min3(
					last,
					row[x],
					row[x+1],
				)
			}
		}
	}
	return row[n]
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}
