package main

import "fmt"

func main() {
	var N, count uint8
	fmt.Scan(&N)
	for i := N; i > 0; i-- {
		var elem int
		fmt.Scan(&elem)
		if elem == 0 {
			count++
		}
	}
	fmt.Print(count)
}
