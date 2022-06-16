package main

import (
	"fmt"
	"time"
)

func goo() {

	m := make(map[int]string, 10)
	m[1] = "abc"
	if value, exists := m[2]; exists {
		fmt.Printf("[%s]\n", value)
	} else if 3 > 9 {
		value = "23"
		exists = false
	} else {
		value = "23"
		exists = false
	}
}

func if_nest() {
	now := time.Now()
	weekday := now.Weekday()
	hour := now.Hour()
	fmt.Println(weekday, hour)
	if weekday.String() != "Sat" && weekday.String() != "Sun" {
		if hour <= 9 && hour >= 7 {
			fmt.Println("Bus lane onlly")
		} else {
			fmt.Println("All vehicle")
		}
	} else {
		fmt.Println("All vehicle")
	}

}

func main() {
	if_nest()
}
