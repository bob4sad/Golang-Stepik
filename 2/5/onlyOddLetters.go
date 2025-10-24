package main

import "fmt"

func main() {
	var X string
	fmt.Scan(&X)

	for i, l := range X {
		if i%2 != 0 {
			fmt.Print(string(l))
		}
	}
}
