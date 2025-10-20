package main

import "fmt"

func main() {
	var N int
	var revBinary []int

	fmt.Scan(&N)

	for {
		if N/2 != 0 {
			revBinary = append(revBinary, N%2)
			N /= 2
		} else {
			revBinary = append(revBinary, N%2)
			break
		}
	}

	for i := len(revBinary) - 1; i >= 0; i-- {
		fmt.Print(revBinary[i])
	}
}
