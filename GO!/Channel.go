package main

import "fmt"

func main() {
	var ch chan int
	//var [name] chan int declaration of channel
	//fmt.Printf("Channel is nil %t\n", ch == nil)
	//fmt.Printf("Len of ch is %d\n", len(ch))

	ch = make(chan int, 10) //10 is capacity
	fmt.Printf("cap of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))

	for i := 0; i < 10; i++ {
		ch <- 3
	}

	fmt.Printf("len of ch is %d\n", len(ch))

	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))
	close(ch)
	for ele := range ch {
		fmt.Println(ele)
	}
}
