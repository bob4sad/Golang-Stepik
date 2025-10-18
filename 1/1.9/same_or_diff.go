package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	switch {
	case a[0] != a[1] && a[1] != a[2]:
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}
