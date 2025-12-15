package _118_Pascals_Triangle

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	result := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, row)
	}

	return result
}
