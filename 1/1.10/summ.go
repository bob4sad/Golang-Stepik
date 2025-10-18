package main

import "fmt"

func main() {
	var a, b, r uint8

	fmt.Scan(&a, &b)

	for a <= b {
		r = r + a
		a++
	}

	fmt.Print(r)
}
