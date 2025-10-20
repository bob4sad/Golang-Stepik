package main

import "fmt"

func main() {
	var N, count, min int
	fmt.Scan(&N)

	if N != 0 {
		fmt.Scan(&min)
		count = 1
		N -= 1
	}

	for i := 0; i < N; i++ {
		var elem int
		fmt.Scan(&elem)
		if elem < min {
			min = elem
			count = 1
		} else if elem == min {
			count++
		}
	}
	fmt.Print(count)
}
