package main

import "fmt"

func main() {
	var x, p, y int
	var count uint8
	fmt.Scan(&x, &p, &y)
	x = x * 100
	y = y * 100
	for x < y {
		x = x + x/p
		count++
	}
	fmt.Print(count)
}
