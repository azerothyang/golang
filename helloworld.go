package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello goroutine")
	}()
	fmt.Println("hello world!")
	time.Sleep(1000 * time.Millisecond)
}
