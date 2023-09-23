package others

func TransposeMatrix(matrix [][]int) (output [][]int) {
	for i := 0; i < len(matrix[0]); i++ {
		o := make([]int, 0)
		o = append(o, matrix[0][i])
		output = append(output, o)
	}
	for k := 0; k < len(output); k++ {
		for i := 1; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if k == j {
					output[k] = append(output[k], matrix[i][j])
				}
			}
		}
	}
	return output
}

func TransposeMatrix2(matrix [][]int) (output [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				output = append(output, make([]int, 0))
			}
			output[j] = append(output[j], matrix[i][j])
		}
	}
	return output
}
