package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n/10 != 1 && n%10 == 1 {
		fmt.Printf("%d korova", n)
	} else if n/10 != 1 && (n%10 == 2 || n%10 == 3 || n%10 == 4) {
		fmt.Printf("%d korovy", n)
	} else {
		fmt.Printf("%d korov", n)
	}
}
