package main

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
