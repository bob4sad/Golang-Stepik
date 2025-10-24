package main

import (
	"fmt"
	"strings"
)

func main() {
	var X, Y string
	fmt.Scan(&X)

	for _, r := range X {
		Y += string(r) + "*"
	}

	fmt.Print(strings.TrimSuffix(Y, "*"))
}
