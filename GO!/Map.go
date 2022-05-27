package main

import "fmt"

func test2() {
	var m2 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	fmt.Printf("m2: %v\n", m2)
}

func test1() {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	fmt.Printf("m1: %v\n", m1)
	//var m1 map[string]string
	//m1 = make(map[string]string)
}

func main() {
	test1()
	test2()
}
