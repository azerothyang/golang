package main

import (
	"time"
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	//arr := strings.Split("dasd", "|")
	//for i, v := range arr {
	//	if v == "" {
	//		continue
	//	}
	//	fmt.Println(i,v)
	//}
	//Jan 2, 2006 at 3:04pm (MST)
	// M D y, HH:MM:SS
	str := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(str)
}
