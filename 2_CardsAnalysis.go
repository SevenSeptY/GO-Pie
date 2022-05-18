/*
Project:
Cards:
newDeck: Create a list of playting cards. Essentially an array of strings
print: Log out the contents of a deck of cards
Shuffle: shuffle all the cards in a deck 
deal : create a hand of cards
saveToFile: Save a list of cards to a file on the local machine 
newDeckFromFile: Load a list of cards from the local machine 
*/
package main

import "fmt"

func main() {
	/*creating new variable and assign value to it
	":=" initialized when first define
	card := "Ace of Spades" / var card string = "Ace of Spades"
	*/
	card := newCard()
	fmt.Println(card)
}
// What data type need to be returned, when newCard called, string retrurn 
func newCard() string {
	return "Five of Diamonds"
}
