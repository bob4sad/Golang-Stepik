package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var slice = make([]int, N, N)

	for i := 0; i < N; i++ {
		var elem int
		fmt.Scan(&elem)
		slice[i] = elem
	}

	fmt.Print(slice[3])
}
