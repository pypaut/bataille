package game

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() (err error) {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("game done")
	}

	g.UpdateMouseStates()

	if g.ShowCards && g.JustClicked && g.IsCursorOnButton {
		g.ShowCards = false

		if g.CurrentWin == 1 {
			g.GiveCardsToPlayerOne()

		} else if g.CurrentWin == -1 {
			g.GiveCardsToPlayerTwo()

		} else if g.CurrentWin == 0 {
			if err = g.CheckCardsAreDuplicates(); err != nil {
				return err
			}

			g.MoveFrontCardsToBack()
		}

		g.ResetWinBadges()

	} else if g.JustClicked && g.IsCursorOnButton {
		g.ShowCards = true

		if err = g.UpdateCurrentWin(); err != nil {
			return err
		}

		g.UpdateWinBadges()
	}

	err = g.CheckTotalNumberOfCards()
	if err != nil {
		return err
	}

	g.UpdatePlayersMessages()
	g.JustClicked = false

	return nil
}

func (g *Game) UpdatePlayersMessages() {
	g.PlayerOneMessage = fmt.Sprintf("%d", len(g.DeckOne.Cards))
	g.PlayerTwoMessage = fmt.Sprintf("%d", len(g.DeckTwo.Cards))
}

func (g *Game) UpdateCurrentWin() error {
	winValue, err := g.DeckOne.WinsAgainst(g.DeckTwo)
	if err != nil {
		return err
	}

	g.CurrentWin = winValue
	return nil
}

func (g *Game) UpdateWinBadges() {
	if g.CurrentWin == 1 {
		g.PlayerOneWins = "+"
		g.PlayerTwoWins = ""
	} else if g.CurrentWin == -1 {
		g.PlayerTwoWins = "+"
		g.PlayerOneWins = ""
	} else {
		g.PlayerTwoWins = ""
		g.PlayerOneWins = ""
	}
}

func (g *Game) ResetWinBadges() {
	g.PlayerTwoWins = ""
	g.PlayerOneWins = ""
}

func (g *Game) GiveCardsToPlayerOne() {
	// Add cards to back of DeckOne
	g.DeckOne.Cards = append(
		g.DeckOne.Cards,
		g.DeckOne.Cards[0],
		g.DeckTwo.Cards[0],
	)

	// Remove cards from both
	g.DeckOne.Cards = g.DeckOne.Cards[1:]
	g.DeckTwo.Cards = g.DeckTwo.Cards[1:]
}

func (g *Game) GiveCardsToPlayerTwo() {
	// Add cards to back of DeckTwo
	g.DeckTwo.Cards = append(
		g.DeckTwo.Cards,
		g.DeckTwo.Cards[0],
		g.DeckOne.Cards[0],
	)

	// Remove cards from both
	g.DeckOne.Cards = g.DeckOne.Cards[1:]
	g.DeckTwo.Cards = g.DeckTwo.Cards[1:]
}

func (g *Game) MoveFrontCardsToBack() {
	// Duplicate front cards to back
	g.DeckOne.Cards = append(g.DeckOne.Cards, g.DeckOne.Cards[0])
	g.DeckTwo.Cards = append(g.DeckTwo.Cards, g.DeckTwo.Cards[0])

	// Remove front cards
	g.DeckOne.Cards = g.DeckOne.Cards[1:]
	g.DeckTwo.Cards = g.DeckTwo.Cards[1:]
}

func (g *Game) CheckCardsAreDuplicates() error {
	if g.DeckOne.Cards[0].Color == g.DeckTwo.Cards[0].Color {
		c := g.DeckOne.Cards[0]
		errorMsg := fmt.Sprintf("duplicate card: %d, %d", c.Color, c.Value)
		return errors.New(errorMsg)
	}

	return nil
}

func (g *Game) CheckTotalNumberOfCards() error {
	if len(g.DeckOne.Cards)+len(g.DeckTwo.Cards) != 52 {
		errorMsg := fmt.Sprintf(
			"game should contain 52 cards (D1: %d, D2: %d)",
			len(g.DeckOne.Cards),
			len(g.DeckTwo.Cards),
		)

		return errors.New(errorMsg)
	}

	return nil
}

func (g *Game) UpdateMouseStates() {
	g.CursorX, g.CursorY = ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.JustClicked = true
		g.Clicking = true
	} else {
		g.JustClicked = false
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.Clicking = false
	}

	if int(g.PlayButtonX) <= g.CursorX &&
		g.CursorX <= int(g.PlayButtonX)+g.PlayButtonImage.Bounds().Dx() &&
		int(g.PlayButtonY) <= g.CursorY &&
		g.CursorY <= int(g.PlayButtonY)+g.PlayButtonImage.Bounds().Dy() {
		g.IsCursorOnButton = true
	} else {
		g.IsCursorOnButton = false
	}
}
