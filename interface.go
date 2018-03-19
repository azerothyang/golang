package main

import "fmt"

type person interface {
	say()
}

type userr struct {
	name string
	age int
	words string
}

func main() {
	var u userr
	u.name = "yang"
	u.age = 1000
	u.words = "hello world"
	speak(&u)
}

func speak(p person)  {
	p.say()
}

func (u *userr)say()  {
	fmt.Println(u.name, u.age, u.words)
}
