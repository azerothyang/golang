package main

import (
	"fmt"
	"time"
)

func main() {
	//mp := sync.Map{}
	//n := 10000
	//for i:=0; i<=n; i++ {
	//	mp.Store(i, rand.Int())
	//}
	//fmt.Println(mp.Load(20))
	s := time.Now().UnixNano()
	fmt.Println(s)
}
