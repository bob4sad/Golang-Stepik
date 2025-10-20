package main

import "fmt"

func main() {
	var N uint8
	fmt.Scan(&N)
	var i uint8 = 0
	for ; i < N; i++ {
		var elem int
		fmt.Scan(&elem)
		if i%2 == 0 {
			fmt.Printf("%d ", elem)
		}
	}

}
