package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)

	text = strings.Trim(text, "\n\t\v\f\r")

	if unicode.IsUpper(rs[0]) && strings.HasSuffix(text, ".") {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}

}
