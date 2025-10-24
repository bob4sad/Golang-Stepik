package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.ToLower(strings.Trim(text, "\f\n\r\t\v"))

	var rs []rune = []rune(text)

	var a, b int = 0, len(rs) - 1
	var status string = "Палиндром"

	for a <= b {
		if rs[a] == rs[b] {
			a++
			b--
		} else {
			status = "Нет"
			break
		}
	}

	fmt.Print(status)
}
