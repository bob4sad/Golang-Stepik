package main

import "fmt"

func powerOf2(x int) int {
	var result int = 1
	for i := 0; i < x; i++ {
		result *= 2
	}
	return result
}

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; N >= powerOf2(i); i++ {
		fmt.Printf("%d ", powerOf2(i))
	}
}
