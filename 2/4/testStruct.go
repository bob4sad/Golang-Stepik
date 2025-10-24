package main

import "fmt"

type TestStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (t *TestStruct) Shoot() bool {
	if t.On == true && t.Ammo > 0 {
		t.Ammo--
		return true
	} else {
		return false
	}
}

func (t *TestStruct) RideBike() bool {
	if t.On == true && t.Power > 0 {
		t.Power--
		return true
	} else {
		return false
	}
}

func main() {

	testStruct := TestStruct{true, 2, 1}

	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.Shoot())

}
