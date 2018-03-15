package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		Cyan = iota
		Megenta
		Yellow
	)
	fmt.Println(Cyan, Megenta, Yellow,  math.Asin(math.Sin(math.Pi/2)))
}
