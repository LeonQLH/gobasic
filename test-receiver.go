package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u User) Notify() {
	fmt.Printf("%v:%v\n", u.Name, u.Email)
}
func testReceiver() {
	u1 := User{"golang", "go@go.com"}
	u1.Notify()
	u2 := User{"gogo", "gogo@gogo.com"}
	u3 := u2
	u3.Notify()
}
