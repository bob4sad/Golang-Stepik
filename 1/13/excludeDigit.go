package main

import "fmt"

func main() {
	var n, d int
	var revResult []int

	fmt.Scan(&n, &d)

	for {
		if n/10 != 0 {
			if n%10 != d {
				revResult = append(revResult, n%10)
			}
		} else {
			if n%10 != d {
				revResult = append(revResult, n%10)
			}
			break
		}
		n /= 10
	}

	for i := len(revResult) - 1; i >= 0; i-- {
		fmt.Print(revResult[i])
	}
}
