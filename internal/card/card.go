package card

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type Color int64

const (
	Heart Color = iota
	Club
	Diamond
	Spade
)

type Value int64

const (
	UnknownValue Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Color Color
	Value Value
	Image *ebiten.Image
}

func ValueFromString(strVal string) Value {
	switch strVal {
	case "2":
		return Two
	case "3":
		return Three
	case "4":
		return Four
	case "5":
		return Five
	case "6":
		return Six
	case "7":
		return Seven
	case "8":
		return Eight
	case "9":
		return Nine
	case "10":
		return Ten
	case "jack":
		return Jack
	case "queen":
		return Queen
	case "king":
		return King
	case "ace":
		return Ace
	}

	return UnknownValue
}

func (c1 *Card) WinsAgainst(c2 *Card) (int, error) {
	if c1.Value == UnknownValue || c2.Value == UnknownValue {
		return -2, errors.New("WinsAgainst: unknown value")
	}

	if c1.Value > c2.Value {
		return 1, nil
	}

	if c1.Value < c2.Value {
		return -1, nil
	}

	// Draw
	return 0, nil
}
