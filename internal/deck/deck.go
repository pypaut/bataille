package deck

import (
	. "bataille/internal/card"

	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Deck struct {
	Cards    []*Card
	BackCard *ebiten.Image
	// Number     Number
	// CardWidth  int
	// CardHeight int
}

func NewDeck() *Deck {
	cards := make([]*Card, 0, 52)

	strValues := []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"jack",
		"queen",
		"king",
		"ace",
	}

	for _, strVal := range strValues {
		// Assign current value
		value := ValueFromString(strVal)

		// Load images
		imgHeart, _, err := ebitenutil.NewImageFromFile(
			fmt.Sprintf("assets/heart_%s.png", strVal),
		)
		if err != nil {
			log.Fatal(err)
		}

		imgClub, _, err := ebitenutil.NewImageFromFile(
			fmt.Sprintf("assets/club_%s.png", strVal),
		)
		if err != nil {
			log.Fatal(err)
		}

		imgDiamond, _, err := ebitenutil.NewImageFromFile(
			fmt.Sprintf("assets/diamond_%s.png", strVal),
		)
		if err != nil {
			log.Fatal(err)
		}

		imgSpade, _, err := ebitenutil.NewImageFromFile(
			fmt.Sprintf("assets/spade_%s.png", strVal),
		)
		if err != nil {
			log.Fatal(err)
		}

		// Create struct cards
		cardHeart := &Card{Color: Heart, Value: value, Image: imgHeart}
		cardClub := &Card{Color: Club, Value: value, Image: imgClub}
		cardDiamond := &Card{Color: Diamond, Value: value, Image: imgDiamond}
		cardSpade := &Card{Color: Spade, Value: value, Image: imgSpade}

		// Add to deck
		cards = append(cards, cardHeart, cardClub, cardDiamond, cardSpade)
	}

	return &Deck{
		Cards: cards,
	}
}

// func (d *Deck) AddCards(cards ...*Card) {
// 	d.Cards = append(d.Cards, cards...)
// }

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) CutInTwo() (*Deck, *Deck) {
	backCardOne, _, err := ebitenutil.NewImageFromFile("assets/back_black.png")
	if err != nil {
		log.Fatal(err)
	}

	deck1 := &Deck{
		Cards:    d.Cards[:26],
		BackCard: backCardOne,
	}

	backCardTwo, _, err := ebitenutil.NewImageFromFile("assets/back_red.png")
	if err != nil {
		log.Fatal(err)
	}

	deck2 := &Deck{
		Cards:    d.Cards[26:],
		BackCard: backCardTwo,
	}

	return deck1, deck2
}

func (d1 *Deck) WinsAgainst(d2 *Deck) (int, error) {
	return d1.Cards[0].WinsAgainst(d2.Cards[0])
}
