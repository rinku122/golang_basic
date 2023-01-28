package main

func main() {
	cards := readFromDrive("myFile")
	cards.shuffle()
	cards.printCards()

	// left, right := afterHand(cards, 2)
	// left.printCards()
	// right.printCards()

	// cards.saveToFile("myFile")

}
