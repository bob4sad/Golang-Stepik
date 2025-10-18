package main

import "fmt"

func main() {
	var a, max, count int

	fmt.Scan(&a)

	for a != 0 {
		if a > max {
			max = a
			count = 1
		} else if a == max {
			count++
		}
		fmt.Scan(&a)
	}

	fmt.Print(count)
}
