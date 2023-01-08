package blackjack

import (
	"fmt"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []Card
}

func (h Hand) value() int {
	result := 0
	numberAces := 0

	for i := 0; i < len(h.Cards); i++ {
		if !strings.ContainsAny(h.Cards[i].Rank, "JQKA") {
			val, _ := strconv.Atoi(h.Cards[i].Rank)
			result += val
		} else {
			switch h.Cards[i].Rank {
			case "J", "Q", "K":
				result += 10
			case "A":
				result += 11
				numberAces += 1
			}
		}
	}
	if result > 21 && numberAces > 1 {
		result -= 10 * numberAces
	}
	return result
}

func (h *Hand) addCard(card Card) {
	h.Cards = append(h.Cards, card)
}

func (h Hand) Display() {
	fmt.Println()
	for _, card := range h.Cards {
		fmt.Print(card.Rank + card.Suit + " ")
	}
}
