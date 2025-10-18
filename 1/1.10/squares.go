package main

import "fmt"

func main() {
	var i uint8 = 1

	for i <= 10 {
		fmt.Println(i * i)
		i++
	}
}
