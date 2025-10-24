package main

import (
	"fmt"
)

func main() {
	var X, max string
	fmt.Scan(&X)

	for _, r := range X {
		if string(r) > max {
			max = string(r)
		}
	}

	fmt.Print(max)
}
