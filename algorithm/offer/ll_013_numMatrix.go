package main

/**
二维子矩阵的和
*/

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{matrix: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	matrix := this.matrix
	for row := row1; row <= row2; row++ {
		for col := col1; col <= col2; col++ {
			sum += matrix[row][col]
		}
	}
	return sum
}

func main() {

}
