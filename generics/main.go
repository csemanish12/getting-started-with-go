package main

import "fmt"

type Number interface {
	int64 | float64
}

func main(){
	// initialize map for integer values
	ints := map[string]int64 {
		"first": 34,
		"second": 12,
	}

	// initialize map for float values
	floats := map[string]float64 {
		"first": 35.98,
		"second": 26.99,
	}
	fmt.Printf("Non-Generic Sums: %v and %v\n",
	 	SumInts(ints), SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
SumIntsOrFloats[string, int64](ints),
SumIntsOrFloats[string, float64](floats))

fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))

fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))

}

func SumInts(m map[string]int64) int64{
	var sum int64
	for _, value := range m {
		sum += value
	}
	return sum
}

func SumFloats(m map[string]float64) float64{
	var sum float64

	for _, value := range m {
		sum += value
	}

	return sum
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}
	return sum
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V

	for _, value := range m {
		sum += value
	}
	return sum
}