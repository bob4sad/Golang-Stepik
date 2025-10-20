package main

import "fmt"

func main() {
	var N uint8
	fmt.Scan(&N)
	var i, count uint8
	var elem int
	for ; i < N; i++ {
		fmt.Scan(&elem)
		if elem > 0 {
			count++
		}
	}

	fmt.Print(count)
}
