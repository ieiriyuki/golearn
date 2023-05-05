package main

import (
	"testing"
)

func TestSumInts(t *testing.T) {
	m := map[string]int64 {
		"a": 1,
		"b": 2,
		"c": 3,
	}
	actual := SumInts(m)
	var expected int64 = 6
	if actual != expected {
		t.Fatalf(`SumInts(m) = %v, want %v`, actual, expected)
	}
}

func TestSumFloats(t *testing.T) {
	m := map[string]float64 {
		"a": 1.1,
		"b": 2.2,
		"c": 3.3,
	}
	actual := SumFloats(m)
	var expected float64 = 6.6
	if actual != expected {
		t.Fatalf(`SumInts(m) = %v, want %v`, actual, expected)
	}
}
