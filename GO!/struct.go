package main

import (
	"fmt"
	"time"
)

type User struct { //can be acccess by other package
	id         int
	Score      float32
	name, addr string
	enrollment time.Time
}

func (user User) hello(man string) string {
	return "hello" + man + ", I am " + user.name
}
func hello(man string, user User) string {
	return "hello" + man + ", I am " + user.name
}

func main() {
	var ws User //user is a strut also a type
	ws = User{Score: 100.00, name: "zxy"}
	//above two rewrite as-> wa := User{Score: 100.00, name: "zxy"}
	ws.Score = 93.5
	ws.enrollment = time.Now()

	a := ws.id + 24
	fmt.Printf("a = %d\n", a)
}
