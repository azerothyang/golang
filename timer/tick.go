package main

import (
	"time"
	"fmt"
)

func main() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v\n", now.Hour())
	}
}
