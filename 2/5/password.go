package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password, result string = "", "Ok"
	fmt.Scan(&password)
	var rs []rune = []rune(password)

	if len(rs) < 5 {
		result = "Wrong password"
	} else {
		for _, r := range rs {
			if !unicode.IsNumber(r) {
				if unicode.Is(unicode.Latin, r) == false {
					result = "Wrong password"
				}
			}
		}
	}

	fmt.Print(result)

}
