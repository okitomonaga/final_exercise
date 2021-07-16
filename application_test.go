package main

import (
	//"reflect"
	//"fmt"
	"testing"
)

func Test_CalcPlus1(t *testing.T) {
	result := calcPlus(3, 6)
	if result != 9 {
		t.Error("Test_CalcPlus1")
	}
}

func Test_CalcPlus2(t *testing.T) {
	result := calcPlus(3, 2)
	if result != 5 {
		t.Error("Test_CalcPlus1")
	}
}

func Test_CalcMinus1(t *testing.T) {
	result := calcMinus(7, 1)
	if result != 6 {
		t.Error("Test_CalcMinus1")
	}
}

func Test_CalcMinus2(t *testing.T) {
	result := calcMinus(5, 1)
	if result != 4 {
		t.Error("Test_CalcMinus2")
	}
}

func Test_CalcMult1(t *testing.T) {
	result := calcMult(6, 7)
	if result != 42 {
		t.Error("Test_CalcMult1")
	}
}
func Test_CalcMult2(t *testing.T) {
	result := calcMult(9, 7)
	if result != 63 {
		t.Error("Test_CalcMult2")
	}
}
