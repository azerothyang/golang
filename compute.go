package main

import (
	"fmt"
)

func main() {
	var n int64 = 1000000000
	times := 1
	for true {
		sum(n)
		fmt.Println(times)
		times++
	}
}

func sum(n int64) int64 {
	var i,res int64
	res = 0
	for i=0; i<=n; i++{
		res += i
	}
	return res
}
