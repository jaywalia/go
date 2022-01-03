package main

import "fmt"

func PrintMatrix(a [][]int) {
	for _, rows := range a {
		for _, elem := range rows {
			fmt.Printf("%d ", elem)
		}
		fmt.Printf("\n")
	}
}

func NewMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for r := range m {
		m[r] = make([]int, cols)
	}
	return m
}

func main() {

	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{5, 6}, {7, 8}}
	C := NewMatrix(2, 2)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				C[i][k] += A[i][j] * B[j][k]

			}
		}
	}
	PrintMatrix(A[:])
	PrintMatrix(B[:])
	PrintMatrix(C[:])

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Printf("%d\t", C[i][j])
	// 	}
	// 	fmt.Println()
	// }
}
