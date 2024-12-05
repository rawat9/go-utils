package matrix

// Transpose returns the transpose of a matrix.
func Transpose[T any](matrix [][]T) [][]T {
	m := len(matrix)
	n := len(matrix[0])

	newMatrix := make([][]T, n)

	for i := 0; i < m; i++ {
		newMatrix[i] = make([]T, m)
		for j := 0; j < n; j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}

	return newMatrix
}
