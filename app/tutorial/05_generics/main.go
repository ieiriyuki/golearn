package main

import "fmt"

func main() {
	ints := map[string]int64{
		"1": 34,
		"2": 12,
	}
	floats := map[string]float64{
		"1": 35.98,
		"2": 26.99,
	}
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)
}
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
