package main

import "fmt"

// function_args
//
func sum(a int, b int) int {
	return a + b
}

func f1(a int) {
	a = 100
}

func f2(s []int) {
	s[0] = 1000
}

func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v:%v\n", v)
	}
}

func f4(name string, ok bool, args ...int) {
	fmt.Println("name:", name)
}

func main() {
	f3(1, 2, 3, 4)
	fmt.Printf("r:%v\n", f3)
	/*x := 200
	f1(x)
	fmt.Printf("r:%v\n", x)
	//r := sum(1, 2)
	//fmt.Printf("r:%v\n", r)
	*/
}
