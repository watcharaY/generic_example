package main

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

// T      -> refer to Number interface
// numbers -> refer to array of any type within Number interface
func SumGeneric[T Number](numbers []T) T {
	var result T
	for _, e := range numbers {
		result += e
	}
	return result
}

// same way as declare below
func SumGeneric2[T int | int32 | int64 | float32 | float64](numbers []T) T {
	var result T
	for _, e := range numbers {
		result += e
	}
	return result
}

// we can use Ordered to handle things for us
// constraints.Ordered can be use in version >= 1.18
func SumGeneric3[T constraints.Ordered](numbers []T) T {
	var result T
	for _, e := range numbers {
		result += e
	}
	return result
}

// cmp.Ordered (standard lib) can be use in version >= 1.21
func SumGeneric4[T cmp.Ordered](numbers []T) T {
	var result T
	for _, e := range numbers {
		result += e
	}
	return result
}
