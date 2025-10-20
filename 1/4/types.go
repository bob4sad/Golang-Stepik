package main

import "fmt"

func main() {
	var b uint8 = 255
	b++ // попробовал переполнить 8 бит
	fmt.Println(b)

	var t float32 = 1.01 - 0.99
	fmt.Println(t)

	var c complex64 = 1 + 2i
	var f complex64 = 1 - 3i
	fmt.Println(c + f)

	fmt.Println(string("Hello!"[0]))   // нормально отрабатывает с англ языком
	fmt.Println(string("Привет!"[0]))  // взят только первый байт, а в UTF-8 русские символы состоят из 2 байт
	fmt.Println(string("Привет!"[:2])) // если взять первые 2 байта, то на выходе получится нормальный русский символ

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}
