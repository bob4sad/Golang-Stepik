package main

import "fmt"

func main() {
	var a, b, max int
	fmt.Scan(&a, &b)

	max = a

	for ; b >= a; b-- {
		if b%7 == 0 && b > max {
			max = b
		}
	}

	if max%7 == 0 {
		fmt.Print(max)
	} else {
		fmt.Print("NO")
	}
}
