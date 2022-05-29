package main

import "fmt"

func test1() {
	m1 := map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	for k := range m1 {
		fmt.Printf("k: %v\n", k)
	}

	for k, v := range m1 {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func main() {
	test1()
}
