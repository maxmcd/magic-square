package main

import "testing"

func isMagicSquare(solution [][]int) bool {
	n := len(solution)
	total := n * (n*n + 1) / 2
	for _, row := range solution {
		var testTotal int
		for _, n := range row {
			testTotal += n
		}
		if testTotal != total {
			return false
		}
	}
	for i := 0; i < n; i++ {
		var testTotal int
		for _, row := range solution {
			testTotal += row[i]
		}
		if testTotal != total {
			return false
		}
	}
	return true
}

func TestSiamese(t *testing.T) {
	numbers := []int{3, 5, 9, 11, 19, 21, 101, 301, 501}
	for _, i := range numbers {
		output := Solve(i)
		if isMagicSquare(output) != true {
			t.Errorf("Not a magic square for %d", i)
		}
	}
}

func TestDoubleEven(t *testing.T) {
	numbers := []int{4, 8, 12, 16, 20, 40, 80, 200}
	for _, i := range numbers {
		output := Solve(i)
		if isMagicSquare(output) != true {
			t.Errorf("Not a magic square for %d", i)
		}
	}
}

func BenchmarkSiamese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(101)
	}
}
func BenchmarkDoubleEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(100)
	}
}

// 1 0 0 1
// 0 1 1 0
// 0 1 1 0
// 1 0 0 1

// 1 0 0 1 1 0 0 1
// 0 1 1 0 0 1 1 0
// 0 1 1 0 0 1 1 0
// 1 0 0 1 1 0 0 1
// 1 0 0 1 1 0 0 1
// 0 1 1 0 0 1 1 0
// 0 1 1 0 0 1 1 0
// 1 0 0 1 1 0 0 1

// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0
// 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1
