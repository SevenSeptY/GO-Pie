/*
Array has fixed 
Slice can be expanded or shirnk 

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
