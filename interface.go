package main

import "fmt"

type admin struct {
	role string
	userr
}

type person interface {
	say()
}

type userr struct {
	name  string
	age   int
	words string
}

func main() {
	var admin admin
	admin.age = 10
	admin.name = "admin"
	admin.role = "admin"
	admin.words = "dasdad"
	var u userr
	u.name = "yang"
	u.age = 1000
	u.words = "hello world"
	speak(&u)
	speak(&admin)
}

func speak(p person) {
	p.say()
}

func (u *userr) say() {
	fmt.Println(u.name, u.age, u.words)
}
