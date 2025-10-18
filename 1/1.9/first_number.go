package main

import "fmt"

func main() {
	var a uint
	fmt.Scan(&a)
	switch {
	case a/10000 > 0:
		fmt.Print(a / 10000)
	case a/1000 > 0:
		fmt.Print(a / 1000)
	case a/100 > 0:
		fmt.Print(a / 100)
	case a/10 > 0:
		fmt.Print(a / 10)
	default:
		fmt.Print(a)
	}
}
