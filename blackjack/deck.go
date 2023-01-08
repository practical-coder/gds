package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) dealCard() Card {
	result := d.Cards[0]
	d.Cards = d.Cards[1:]
	return result
}

func (d *Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) initialize() Deck {
	for _, suit := range suits {
		for _, rank := range ranks {
			d.Cards = append(d.Cards, Card{rank, string(suit)})
		}
	}
	d.shuffle()
	return *d
}

//lint:ignore U1000 displaying full deck is not used for the time being
func (d Deck) display() {
	for _, card := range d.Cards {
		fmt.Printf("%s%s ", card.Rank, card.Suit)
	}
}
