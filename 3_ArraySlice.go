/*
Array has fixed 
Slice can be expanded or shirnk 
1. With the var keyword:
var array_name = [length]datatype{values} // here length is defined
or
var array_name = [...]datatype{values} // here length is inferred

2. With the := sign:
array_name := [length]datatype{values} // here length is defined
or
array_name := [...]datatype{values} // here length is inferred
*/
package main
import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
//
  for i, card := range cards {//range cards: Take the slice of 'cards' and loop over it 
		fmt.Println(i, card)      //Run this one time for each card in the slice
	}
}

// What data type need to be returned, when newCard called, string retrurn
func newCard() string {
	return "Five of Diamonds"
}

//------------------------Array example--------------
package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
