package main

import "fmt"

const (
	test   = "нежданчик"
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	test2    = "нежданчик"
	Thursday = iota
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Saturday)
}
