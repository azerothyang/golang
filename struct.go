package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func main() {
	var u user
	u.name = "yang"
	u.age = 100
	u.changeAge(10)
	fmt.Println(u)
}

func (user user) getName() string {
	return user.name
}

func (puser *user)changeAge(age uint8) {
	puser.age = age
}
