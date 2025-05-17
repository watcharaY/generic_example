package main

import (
	"testing"
)

func Test_SumGeneric_Int(t *testing.T) {
	intArray := []int{1, 2, 3, 4, 5}
	wantInt := 15
	result := SumGeneric(intArray)
	if result != wantInt {
		t.Errorf("given %v want %v but got %v", intArray, wantInt, result)
	}
}
func Test_SumGeneric_Int32(t *testing.T) {
	int32Array := []int32{1, 2, 3, 4, 5}
	wantInt32 := int32(15)
	result := SumGeneric(int32Array)
	if result != wantInt32 {
		t.Errorf("given %v want %v but got %v", int32Array, wantInt32, result)
	}
}
func Test_SumGeneric_Int64(t *testing.T) {
	int64Array := []int64{1, 2, 3, 4, 5}
	wantInt64 := int64(15)
	result := SumGeneric(int64Array)
	if result != wantInt64 {
		t.Errorf("given %v want %v but got %v", int64Array, wantInt64, result)
	}
}
func Test_SumGeneric_Float32(t *testing.T) {
	float32Array := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	wantFloat32 := float32(16.5)
	result := SumGeneric(float32Array)
	if result != wantFloat32 {
		t.Errorf("given %v want %v but got %v", float32Array, wantFloat32, result)
	}
}
func Test_SumGeneric_Float64(t *testing.T) {
	float64Array := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	wantFloat64 := float64(16.5)
	result := SumGeneric(float64Array)
	if result != wantFloat64 {
		t.Errorf("given %v want %v but got %v", float64Array, wantFloat64, result)
	}
}
