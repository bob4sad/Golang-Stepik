package main

import "fmt"

func replaceValue(x1 *int, x2 *int) {
	*x1 += *x2
	*x2 = *x1 - *x2
	*x1 -= *x2
}

func main() {
	var a, b int = 2, 3
	replaceValue(&a, &b)
	fmt.Println(a, b)
}
