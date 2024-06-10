package main

import "fmt"

type User struct {
	Age []int8
}

func (u User) ShowAge() {
	fmt.Println(u.Age)
}

func (u *User) ShowAges() {
	fmt.Println(u.Age)
}

func (u *User) AddAges(age int8) {
	u.Age = append(u.Age, age)
}

func createUser1() User {
	return User{
		Age: []int8{1},
	}
}

func createUser2() *User {
	return &User{
		Age: []int8{1},
	}
}
