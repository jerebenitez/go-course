package L13

func createMatrix(rows, cols int) [][]int {
    matrix := [][]int{}

    for y := 0; y < rows; y++ {
        row := []int{}
        for x := 0; x < cols; x++ {
            row = append(row, y * x)
        }
        matrix = append(matrix, row)
    }

    return matrix
}

