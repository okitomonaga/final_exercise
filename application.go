package main

import "fmt"
import "strconv"
import "math"
import "gonum.org/v1/gonum/mat"

func PrintCoordinate(X mat.Matrix) string {
	return "(" + strconv.FormatFloat(X.At(0, 0), 'f', 2, 64) + "," + strconv.FormatFloat(X.At(1, 0), 'f', 2, 64) + ")"

	//return "(1,2)"
}

func parallelTp(X mat.Matrix, x float64, y float64) mat.Matrix {
	tp := []float64{1, 0, x, 0, 1, y, 0, 0, 1}
	tp_mat := mat.NewDense(3, 3, tp)
	result := mat.NewDense(3, 1, nil)
	result.Product(tp_mat, X)
	return result
}

func rotateTp(X mat.Matrix, rad float64) mat.Matrix {
	rotate := []float64{math.Round(math.Cos(rad*math.Pi/180)*100) / 100, -math.Round(math.Sin(rad*math.Pi/180)*100) / 100, 0, math.Round(math.Sin(rad*math.Pi/180)*100) / 100, math.Round(math.Cos(rad*math.Pi/180)*100) / 100, 0, 0, 0, 1}
	rotate_mat := mat.NewDense(3, 3, rotate)
	result := mat.NewDense(3, 1, nil)
	result.Product(rotate_mat, X)
	matPrint(result)
	return result
	//	return mat.NewDense(3, 1, []float64{-2, 1, 1})

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
	fmt.Println(PrintCoordinate(operation_target))
	//matPrint(operation_target)

	// result1 := parallelTp(operation_target, 1.0, 1.0)

	// matPrint(result1)

	fmt.Println("どの操作を行いますか？1. 回転移動, 2. 平行移動, 3. 対称移動, 4. 拡大縮小")
	var mode string
	fmt.Scan(&mode)
	switch mode {
	case "1":
		fmt.Println("回転する大きさを決めてください。")
		var rotate string
		fmt.Scan(rotate)
		rotate, _ := strconv.Atoi(rotate)
		result := rotateTp(operation_target, float64(rotate))
		fmt.Println("最終出力は以下の通りです。")
		matPrint(result)
	case "2":
		fmt.Println("x軸方向、y軸方向の移動量を入力してください。")
		var parallelX_s string
		var parallelY_s string
		fmt.Scan(&parallelX_s, &parallelY_s)
		parallelX, _ := strconv.Atoi(parallelX_s)
		parallelY, _ := strconv.Atoi(parallelY_s)
		result := parallelTp(operation_target, float64(parallelX), float64(parallelY))
		fmt.Printf("最終出力は" + PrintCoordinate(result) + "です。\n")
		//matPrint(result)
	case "3":
		fmt.Println("実装予定です。")
	case "4":
		fmt.Println("x軸、y軸方向の拡大・縮小の倍率を入力してください。")
		var scaleX_s string
		var scaleY_s string
		fmt.Scan(&scaleX_s, &scaleY_s)
		scaleX, _ := strconv.ParseFloat(scaleX_s, 64)
		scaleY, _ := strconv.ParseFloat(scaleY_s, 64)
		result := scaleSlice(operation_target, scaleX, scaleY)
		fmt.Printf("最終出力は" + PrintCoordinate(result) + "です。\n")
		//matPrint(result)
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
