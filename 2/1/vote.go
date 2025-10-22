package main

import "fmt"

func voting(x, y, z uint8) uint8 {
	if x+y+z > 1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Print(voting(0, 0, 1))
}
