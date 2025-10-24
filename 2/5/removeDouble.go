package main

import (
	"fmt"
	"strings"
)

func main() {
	var X, result string
	fmt.Scan(&X)

	for _, l := range X {
		if strings.Count(X, string(l)) == 1 {
			result += string(l)
		}
	}

	fmt.Print(result)
}
