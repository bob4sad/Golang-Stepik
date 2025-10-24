package main

import (
	"fmt"
)

func main() {
	var X int
	var result string
	fmt.Scan(&X)

	for X%10 != 0 {
		result = fmt.Sprint((X%10)*(X%10)) + result
		X /= 10
	}

	fmt.Print(result)
}
