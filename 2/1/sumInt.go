package main

import "fmt"

func sumInt(a ...int) (int, int) {
	var n, summ int
	for _, v := range a {
		n++
		summ += v
	}
	return n, summ
}

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}
