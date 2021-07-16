package main

import (
	//"reflect"
	//"fmt"
	"testing"
)

func TestParallelTp1(t *testing.T) {
	result := ParallelTp("1,2")
	if result[0] != 1 || result[1] != 2 { //resultと"hoge"の型を比較している
		t.Error("ParallelTp1 is failed")
	}
}
