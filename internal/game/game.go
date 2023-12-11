package game

import (
	"bataille/internal/deck"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"os"

	"errors"
	"fmt"

	"github.com/golang/freetype"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	Width  int
	Height int

	CardWidth  int
	CardHeight int

	DeckOne *deck.Deck
	DeckTwo *deck.Deck

	ShowCards     bool
	ScreenMessage string
	DrawOptions   ebiten.DrawImageOptions

	FontFace  font.Face
	FontColor color.Color
}

func NewGame() *Game {
	mainDeck := deck.NewDeck()
	mainDeck.Shuffle()

	deckOne, deckTwo := mainDeck.CutInTwo()
	cardsWidth, cardsHeight := deckOne.Cards[0].Image.Size()
	var drawOptions ebiten.DrawImageOptions

	// Load font
	fontfile := "assets/kongtext.ttf"
	fontBytes, err := os.ReadFile(fontfile)
	if err != nil {
		log.Println(err)
		return nil
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return nil
	}

	fontFace := truetype.NewFace(f, &truetype.Options{Size: 12, DPI: 72, Hinting: font.HintingFull})
	fontColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	return &Game{
		Width:  1280,
		Height: 720,

		CardWidth:  cardsWidth,
		CardHeight: cardsHeight,

		DeckOne: deckOne,
		DeckTwo: deckTwo,

		ShowCards:   false,
		DrawOptions: drawOptions,

		FontFace:  fontFace,
		FontColor: fontColor,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("Game done!")
	}

	justClicked := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)

	if g.ShowCards && justClicked {
		msg := ""

		winValue, err := g.DeckOne.WinsAgainst(g.DeckTwo)
		if err != nil {
			return err
		}

		if winValue == 1 {
			// Add cards to back of DeckOne
			g.DeckOne.Cards = append(
				g.DeckOne.Cards,
				g.DeckOne.Cards[0],
				g.DeckTwo.Cards[0],
			)

			msg = "Player one wins!"

		} else if winValue == -1 {
			// Add cards to back of DeckTwo
			g.DeckTwo.Cards = append(
				g.DeckTwo.Cards,
				g.DeckTwo.Cards[0],
				g.DeckOne.Cards[0],
			)

			msg = "Player two wins!"
		} else {
			// Check if duplicate
			if g.DeckOne.Cards[0].Color == g.DeckTwo.Cards[0].Color {
				c := g.DeckOne.Cards[0]
				errorMsg := fmt.Sprintf("duplicate card: %d, %d", c.Color, c.Value)
				return errors.New(errorMsg)
			}

			// Move both front cards to the back
			g.DeckOne.Cards = append(g.DeckOne.Cards, g.DeckOne.Cards[0])
			g.DeckTwo.Cards = append(g.DeckTwo.Cards, g.DeckTwo.Cards[0])
			msg = "It's a draw!"
		}

		// Rm first cards
		g.DeckOne.Cards = g.DeckOne.Cards[1:]
		g.DeckTwo.Cards = g.DeckTwo.Cards[1:]

		g.ShowCards = false
		g.ScreenMessage = fmt.Sprintf(
			"Deck 1: %d\nDeck 2: %d\n%s", len(g.DeckOne.Cards), len(g.DeckTwo.Cards), msg,
		)
	} else if justClicked {
		g.ShowCards = true
		g.ScreenMessage = fmt.Sprintf(
			"Deck 1: %d\nDeck 2: %d\n", len(g.DeckOne.Cards), len(g.DeckTwo.Cards),
		)
	}

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

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Player 1
	g.DrawOptions.GeoM.Reset()
	g.DrawOptions.GeoM.Translate(
		float64(g.Width-g.CardWidth)/2,
		float64(g.Height-g.CardHeight)/2+float64(g.CardHeight+g.Height/9),
	)
	screen.DrawImage(g.DeckOne.BackCard, &g.DrawOptions)

	if g.ShowCards {
		g.DrawOptions.GeoM.Reset()
		g.DrawOptions.GeoM.Translate(
			float64(g.Width-g.CardWidth)/2,
			float64(g.Height-g.CardHeight)/2+float64(g.CardHeight)/2,
		)

		screen.DrawImage(g.DeckOne.Cards[0].Image, &g.DrawOptions)
	}

	// Player 2
	g.DrawOptions.GeoM.Reset()
	g.DrawOptions.GeoM.Translate(
		float64(g.Width-g.CardWidth)/2,
		float64(g.Height-g.CardHeight)/2-float64(g.CardHeight+g.Height/9),
	)
	screen.DrawImage(g.DeckTwo.BackCard, &g.DrawOptions)

	if g.ShowCards {
		g.DrawOptions.GeoM.Reset()
		g.DrawOptions.GeoM.Translate(
			float64(g.Width-g.CardWidth)/2,
			float64(g.Height-g.CardHeight)/2-float64(g.CardHeight)/2,
		)

		screen.DrawImage(g.DeckTwo.Cards[0].Image, &g.DrawOptions)
	}

	text.Draw(screen, g.ScreenMessage, g.FontFace, g.Width-g.Width/5, g.Height-g.Height/15, g.FontColor)

	return
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
