package main

import "fmt"

func main() {
	workArray := [10]uint8{}

	for i := 0; i < len(workArray); i++ {
		var newElem uint8
		fmt.Scan(&newElem)
		workArray[i] = newElem
	}

	for i := 0; i < 3; i++ {
		var start, finish int
		fmt.Scan(&start, &finish)
		workArray[start] += workArray[finish]
		workArray[finish] = workArray[start] - workArray[finish]
		workArray[start] -= workArray[finish]

	}

	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}

}
