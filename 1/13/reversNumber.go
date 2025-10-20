package main

import "fmt"

func digitByIndex(number int, idx int) int {
	var a, b int = 1, 1
	for i := 0; i < idx; i++ {
		a *= 10
	}
	for i := 0; i < idx-1; i++ {
		b *= 10
	}
	return number % a / b
}

func main() {
	const N int = 3
	var number int
	fmt.Scan(&number)

	for i := 1; i <= N; i++ {
		fmt.Print(digitByIndex(number, i))
	}

}
