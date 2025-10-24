package main

import "fmt"

func summ(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

func main() {
	var a, b int = 2, 2
	test(&a, &b)
}
