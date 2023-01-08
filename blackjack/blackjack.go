package blackjack

import (
	"bufio"
	"fmt"
	"os"
)

var ranks = []string{
	"2", "3", "4", "5", "6", "7", "8",
	"9", "10", "J", "Q", "K", "A",
}

var suits = []rune{'\u2660', '\u2661', '\u2662', '\u2663'}

type Card struct {
	Rank string
	Suit string
}

func Play() {
	gameOver := false
	myDeck := Deck{}
	myDeck.initialize()
	houseHand := Hand{}
	playerHand := Hand{}
	for i := 1; i <= 2; i++ {
		card := myDeck.dealCard()
		houseHand.addCard(card)
		card = myDeck.dealCard()
		playerHand.addCard(card)
	}
	playerHand.Display()

	for {
		fmt.Println("\tDo you want to be hit (y/n)?")
		reader := bufio.NewReader(os.Stdin)
		res, _, _ := reader.ReadRune()
		if res != 'y' {
			break
		}
		card := myDeck.dealCard()
		playerHand.addCard(card)
		playerHand.Display()
		if playerHand.value() > 21 {
			fmt.Println("Player's score exceeds 21. Game over!. House wins!")
			gameOver = true
			break
		}
	}

	if !gameOver {
		for {
			if houseHand.value() > 21 {
				fmt.Println("House score exceeds 21. Game over! Player wins!")
				gameOver = true
				break
			}
			if houseHand.value() < 17 {
				card := myDeck.dealCard()
				houseHand.addCard(card)
			} else {
				break
			}
		}
	}

	if !gameOver {
		if playerHand.value() > houseHand.value() {
			fmt.Println("Player score exceeds House score. Game over! Player wins!")
		} else if playerHand.value() == houseHand.value() {
			fmt.Println("Player score equals House score. Game over! Tie!")
		} else {
			fmt.Println("House score exceeds player score. Game over! House wins!")
		}
	}
}
