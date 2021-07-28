package main

import (
	//"reflect"
	//"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestScaleSlice1(t *testing.T) {
	x := []float64{1.0, 2.0, 1}
	operation_target := mat.NewDense(3, 1, x) //操作対象
	result := scaleSlice(operation_target, 2.0, 3.0)
	if result.At(0, 0) != 2 || result.At(1, 0) != 6 {
		t.Error("TestScaleSlice1 is failed")
	}
}

func TestScaleSlice2(t *testing.T) {
	x := []float64{3.0, 5.0, 1}
	operation_target := mat.NewDense(3, 1, x) //操作対象
	result := scaleSlice(operation_target, 2.0, 3.0)
	if result.At(0, 0) != 6 || result.At(1, 0) != 15 {
		t.Error("TestScaleSlice2 is failed")
	}
}

func TestParallelTp1(t *testing.T) {
	x := []float64{1, 2, 1}
	operation_target := mat.NewDense(3, 1, x)
	result := parallelTp(operation_target, 1, 1)

	if result.At(0, 0) != 2 || result.At(1, 0) != 3 {
		t.Error("ParallelTp1 is failed")
	}
}

func TestParallelTp2(t *testing.T) {
	x := []float64{3, 3, 1}
	operation_target := mat.NewDense(3, 1, x)
	result := parallelTp(operation_target, 1, 1)

	if result.At(0, 0) != 4 || result.At(1, 0) != 4 {
		t.Error("ParallelTp2 is failed")
	}

}

func TestRotateTp1(t *testing.T) {
	x := []float64{1, 2, 1}
	operation_target := mat.NewDense(3, 1, x)
	result := rotateTp(operation_target, 90)

	if result.At(0, 0) != -2 || result.At(1, 0) != 1 {
		t.Error("RotateTp1 is failed")
	}
}

func TestRotateTp2(t *testing.T) {
	x := []float64{2, 3, 1}
	operation_target := mat.NewDense(3, 1, x)
	result := rotateTp(operation_target, 180)

	if result.At(0, 0) != -2 || result.At(1, 0) != -3 {
		t.Error("RotateTp2 is failed")
  }
}

func TestPrintCoordinate1(t *testing.T) {
	x := []float64{1, 2, 1}
	operation_target := mat.NewDense(3, 1, x)
	result := PrintCoordinate(operation_target)

	if result != "(1.00,2.00)" {
		t.Error("TestPrintCoordinate1 is failed")
	}
}

func TestPrintCoordinate2(t *testing.T) {
	x := []float64{3, 4, 1}
	operation_target := mat.NewDense(3, 1, x)
	result := PrintCoordinate(operation_target)

	if result != "(3.00,4.00)" {
		t.Error("TestPrintCoordinate2 is failed")

	}
}
