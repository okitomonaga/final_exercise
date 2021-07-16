package main

import (
	//"reflect"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestParallelTp1(t *testing.T) {
	x := []float64{1, 2, 1}
	operation_target := mat.NewDense(3, 1, x)
	result := parallelTp(operation_target, 1, 1)
	fa := mat.Formatted(result, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)

	if result.At(0, 0) != 2 || result.At(1, 0) != 3 {
		t.Error("ParallelTp1 is failed")
	}
}
