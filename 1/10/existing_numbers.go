package main

import "fmt"

func main() {
	var x, y, y_temp, rev int
	for fmt.Scan(&x, &y); x > 0; x /= 10 {
		for y_temp = y; y_temp > 0; y_temp /= 10 {
			if x%10 == y_temp%10 {
				rev = rev*10 + x%10
			}
		}
	}
	for ; rev > 0; rev /= 10 {
		fmt.Print(rev%10, " ")
	}
}
