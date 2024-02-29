package main

import "fmt"

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