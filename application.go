package main

import "fmt"
import "strconv"
import "gonum.org/v1/gonum/mat"

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func main() {
	var a string
	var b string

	fmt.Scan(&a, &b)
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	x := []float64{ai, bi, 1}
	operation_target := mat.NewDense(3, 1, x) //操作対象
	matPrint(operation_target)

}

// func main() {
//     x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
//     A := mat.NewDense(3, 4, x)
//     // ⎡1  2  3  4⎤
//     // ⎢5  6  7  8⎥
//     // ⎣9 10 11 12⎦
// }
