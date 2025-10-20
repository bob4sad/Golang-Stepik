package main

import "fmt"

func main() {
	var A, previousF, currentF, count int = 0, 1, 2, 3
	fmt.Scan(&A)
	for {
		if A == currentF {
			break
		} else if A < currentF {
			count = -1
			break
		}

		currentF += previousF
		previousF = currentF - previousF
		count++
	}

	fmt.Print(count)

}
