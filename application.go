package main

import "fmt"
import "strconv"
import "gonum.org/v1/gonum/mat"

func parallelTp(X mat.Matrix, x float64, y float64) mat.Matrix {
	return mat.NewDense(3, 1, []float64{2, 3, 1})
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func scaleSlice(target mat.Matrix, a float64, b float64) mat.Matrix {
	// x := []float64{float64(6), float64(3), 1}
	// result := mat.NewDense(3, 1, x) //操作対象
	scale := []float64{a, 0, 0, 0, b, 0, 0, 0, 1}
	scale_mat := mat.NewDense(3, 3, scale)
	result := mat.NewDense(3, 1, nil) //全て0
	result.Product(scale_mat, target)
	return result
}

func main() {
	var a string
	var b string

	fmt.Scan(&a, &b)
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	x := []float64{float64(ai), float64(bi), 1}
	operation_target := mat.NewDense(3, 1, x) //操作対象
	matPrint(operation_target)

	result1 := parallelTp(operation_target, 1.0, 1.0)

	matPrint(result1)
}

// func main() {
//     x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
//     A := mat.NewDense(3, 4, x)
//     // ⎡1  2  3  4⎤
//     // ⎢5  6  7  8⎥
//     // ⎣9 10 11 12⎦
// }
