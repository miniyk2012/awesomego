package main

import "fmt"

// 旋转矩阵, 新申请空间
func rotate(matrix [][]int) {
	newMatrix := make([][]int, len(matrix))
	for i := range newMatrix {
		newMatrix[i] = make([]int, len(matrix))
	}
	for i := range matrix {
		for j := range matrix[i] {
			newMatrix[j][len(matrix)-i-1] = matrix[i][j]
		}
	}
	copy(matrix, newMatrix)
}

// 先对角线翻转, 再水平翻转
func rotateInplace(matrix [][]int) {
	// 先按对角线翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)-i; j++ {
			matrix[len(matrix)-j-1][len(matrix)-i-1], matrix[i][j] = matrix[i][j], matrix[len(matrix)-j-1][len(matrix)-i-1]
		}
	}
	// 再上下翻转
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[len(matrix)-i-1][j], matrix[i][j] = matrix[i][j], matrix[len(matrix)-i-1][j]
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	printMatrix(matrix)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotateInplace(matrix)
	printMatrix(matrix)
}
