package main

import "fmt"

func main() {
	slice := make([]int, 3)
	newSlice := foo(slice)
	fmt.Println(slice)
	fmt.Println(newSlice)

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))
}

func foo(slice []int) []int {
	slice[0] = 100
	return slice
}
