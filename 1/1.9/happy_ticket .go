package main

import "fmt"

func main() {
	var a uint
	fmt.Scan(&a)
	var (
		first  uint = a % 10
		second uint = a % 100 / 10
		third  uint = a % 1000 / 100
		fourth uint = a % 10000 / 1000
		fifth  uint = a % 100000 / 10000
		sixth  uint = a / 100000
	)

	if first+second+third == fourth+fifth+sixth {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
