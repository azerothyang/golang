package main

import "fmt"

func main() {
	testMap := make(map[string]string)
	testMap["key"] = "value"
	for k := range testMap {
		fmt.Println(k)
	}
}
