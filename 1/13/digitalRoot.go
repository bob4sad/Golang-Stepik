package main

import "fmt"

func main() {
	var number, summ int
	fmt.Scan(&number)

	for {
		if number/10 != 0 {
			summ += number % 10
			number /= 10
		} else if (summ+number)/10 != 0 {
			number = summ + number
			summ = 0
		} else {
			summ += number
			break
		}
	}

	fmt.Print(summ)
}
