package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Println(string(a[len(a)-1]))
}
