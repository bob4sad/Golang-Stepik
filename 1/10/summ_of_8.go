package main

import "fmt"

func main() {
	var a, b, r int

	fmt.Scan(&a)

	for ; a != 0; a-- {
		fmt.Scan(&b)
		if b%8 == 0 && b/100 == 0 && b/10 != 0 {
			r += b
		}
	}

	fmt.Print(r)
}
