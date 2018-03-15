package main

import "fmt"

func main() {
	var s = []string {"x", "xx", "xxx"}
	str := "abcd"
	for i, v := range s {
		fmt.Println(i, v)
	}
	for i := range str{
		fmt.Println(i)
	}
}
