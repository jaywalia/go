package main

import (
	"errors"
	"fmt"
	"log"
)

func fastExponent(n int, exp int) int {
	fmt.Printf("n: %v, exp: %v\n", n, exp)
	var r int
	if exp == 1 {
		// base case
		r = n
	} else if exp%2 == 0 {
		// evens
		r = fastExponent(n*n, exp/2)
	} else {
		// odds other than 1
		r = n * fastExponent(n*n, (exp-1)/2)
	}
	return r
}

func printMatrix(a [][]int) {
	for _, rows := range a {
		for _, elem := range rows {
			fmt.Printf("%d ", elem)
		}
		fmt.Printf("\n")
	}
}

func newMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for r := range m {
		m[r] = make([]int, cols)
	}
	return m
}

func matrixMul(A [][]int, B [][]int) ([][]int, error) {
	A_rows := len(A)
	A_cols := len(A[0]) // A_cols or B_rows
	B_rows := len(B)
	if A_cols != B_rows {
		return nil,
			errors.New("Cannot multiply A*B as cols of A is not same as rows of B")
	}
	B_cols := len(B[0])
	C := newMatrix(A_rows, B_cols)
	for i := 0; i < A_rows; i++ {
		for j := 0; j < A_cols; j++ {
			for k := 0; k < B_cols; k++ {
				C[i][k] += A[i][j] * B[j][k]
			}
		}
	}
	return C, nil
}

func matrixMulMod(A [][]int, B [][]int, mod int) ([][]int, error) {
	A_rows := len(A)
	A_cols := len(A[0]) // A_cols or B_rows
	B_rows := len(B)
	if A_cols != B_rows {
		return nil,
			errors.New("Cannot multiply A*B as cols of A is not same as rows of B")
	}
	B_cols := len(B[0])
	C := newMatrix(A_rows, B_cols)
	for i := 0; i < A_rows; i++ {
		for j := 0; j < A_cols; j++ {
			for k := 0; k < B_cols; k++ {
				ab := (A[i][j] * B[j][k]) % mod
				C[i][k] = (C[i][k] + ab) % mod
			}
		}
	}
	return C, nil
}

func fibIteration(a int, b int, n int) int {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55,
	// 89, 144, 233, 377, 610, 987, 1597, 2584,
	// 4181, 6765, 10946, 17711, 28657, 46368,
	// 75025, 121393, 196418, 317811

	var f0 int = a
	var f1 int = b
	for n > 0 {
		// fmt.Println(f0)
		f0, f1 = f1, (f0 + f1)
		n--
	}
	return f0
}

func fibIterationMod(a int, b int, n int, m int) int {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55,
	// 89, 144, 233, 377, 610, 987, 1597, 2584,
	// 4181, 6765, 10946, 17711, 28657, 46368,
	// 75025, 121393, 196418, 317811

	var f0 int = a
	var f1 int = b
	for n > 0 {
		// fmt.Println(f0)
		f0, f1 = f1, (f0+f1)%m
		n--
	}
	return f0
}

func getFibInitialMatrix() [][]int {
	return [][]int{{1, 1}, {1, 0}}
}

func getVector(a int, b int) [][]int {
	return [][]int{{a}, {b}}
}

func matrixMulFastExponent(M [][]int, exp int, mod int) [][]int {
	var R [][]int
	M2, error := matrixMulMod(M, M, mod)
	if error != nil {
		log.Fatal(error)
	}

	if exp == 1 {
		R = M
	} else if exp%2 == 0 {
		R = matrixMulFastExponent(M2, exp/2, mod)
	} else {
		MEven := matrixMulFastExponent(M2, (exp-1)/2, mod)
		R, error = matrixMulMod(M, MEven, mod)
		if error != nil {
			log.Fatal(error)
		}
	}
	return R
}

func fibMatrixMod(a int, b int, n int, m int) int {
	// fib matrix
	// fast exponentian
	// matrix * vector multiply
	// get answer

	// var m int = 1000*1000*1000 + 7 // prime for modulus
	F := getFibInitialMatrix()
	F2n := matrixMulFastExponent(F, n, m)
	V := getVector(a, b)
	F2nV, error := matrixMulMod(F2n, V, m)
	if error != nil {
		log.Fatal(error)
	}

	return F2nV[0][0]
}

func main() {
	//r := fastExponent(2, 10)
	//fmt.Println(r)

	//509618737 460201239 229176339
	var a int = 0
	var b int = 1
	// var a int = 460201239
	// var b int = 509618737
	// var n int = 229176339
	var n int = 10
	var m int = 1000*1000*1000 + 7 // a large prime
	// fmt.Println(n % m)
	// var n int = 20
	fm := fibMatrixMod(a, b, n, m)
	fi := fibIterationMod(a, b, n, m)
	fmt.Printf("fm: %v, fi:%v \n", fm, fi)
}
