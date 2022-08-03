package main

import "fmt"

type notifier interface {
	notify()
}

//define a struct admin
type admin struct {
	name string
}

//type admin struct {
//	level int
//	user
//}

func (a *admin) notify() {
	fmt.Printf("admin %s notified!\n", a.name)
}

//define a struct user
type user struct {
	//fields
	name  string
	email string
}

//
func (u user) notify() {
	fmt.Printf("user %s notified!\n", u.name)
}

//
func (u user) printName() {
	fmt.Println(u.name)
}

// accept the point
func (u *user) changeName(newName string) {
	u.name = newName
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{
		name:  "young",
		email: "young@163.com",
	}

	u.printName()
	u.notify()
	u.changeName("Wang")
	u.printName()

	sendNotification(u)
	sendNotification(&admin{name: "xiaoming"})

}
