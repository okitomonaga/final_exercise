package main

import "fmt"
import "strconv"
import "gonum.org/v1/gonum/mat"

func parallelTp(X mat.Matrix, x float64, y float64) mat.Matrix {
	tp := []float64{1, 0, x, 0, 1, y, 0, 0, 1}
	tp_mat := mat.NewDense(3, 3, tp)
	result := mat.NewDense(3, 1, nil)
	result.MulElem(tp_mat, X)
	return result
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

	fmt.Println("操作する座標を入力してください")
	fmt.Scan(&a, &b)
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	x := []float64{float64(ai), float64(bi), 1}
	operation_target := mat.NewDense(3, 1, x) //操作対象
	matPrint(operation_target)

	// result1 := parallelTp(operation_target, 1.0, 1.0)

	// matPrint(result1)

	fmt.Println("どの操作を行いますか？1. 回転移動, 2. 平行移動, 3. 対称移動, 4. 拡大縮小")
	var mode string
	fmt.Scan(&mode)
	switch mode {
	case "1":
	case "2":
		fmt.Println("x軸方向、y軸方向の移動量を入力してください。")
		var parallelX_s string
		var parallelY_s string
		fmt.Scan(&parallelX_s, &parallelY_s)
		parallelX, _ := strconv.Atoi(parallelX_s)
		parallelY, _ := strconv.Atoi(parallelY_s)
		result := parallelTp(operation_target, float64(parallelX), float64(parallelY))
		fmt.Println("最終出力は以下の通りです。")
		matPrint(result)
	case "3":
	case "4":
		fmt.Println("x軸、y軸方向の拡大・縮小の倍率を入力してください。")
		var scaleX_s string
		var scaleY_s string
		fmt.Scan(&scaleX_s, &scaleY_s)
		scaleX, _ := strconv.Atoi(scaleX_s)
		scaleY, _ := strconv.Atoi(scaleY_s)
		result := scaleSlice(operation_target, float64(scaleX), float64(scaleY))
		fmt.Println(scaleX)
		fmt.Println(scaleY)
		fmt.Println("最終出力は以下の通りです。")
		matPrint(result)
	default:
		fmt.Println("不正な入力です")
	}
}

// func main() {
//     x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
//     A := mat.NewDense(3, 4, x)
//     // ⎡1  2  3  4⎤
//     // ⎢5  6  7  8⎥
//     // ⎣9 10 11 12⎦
// }
