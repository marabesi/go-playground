package main

import "fmt"


func main() {
	fmt.Println(newInt() == newDummyInt())

	fmt.Printf("%d\n", newInt())
	fmt.Printf("%d\n", newInt())

	fmt.Printf("========================\n")

    randomInt := newInt()

    fmt.Printf("%d\n", randomInt)
	fmt.Printf("%d\n", randomInt)
}

func newInt() *int {
	return new(int)
}

func newDummyInt() *int {
	var dummy int
	return &dummy
}