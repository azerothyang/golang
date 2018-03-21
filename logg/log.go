package main

import (
	"log"
	"fmt"
)

func init()  {
	log.SetPrefix("TRACE：")
	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Llongfile)
}

func main() {
	log.Println("hello")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println(222)
	}()
	panic(55)
	log.Println(111)
}
