package game

import (
	"bataille/internal/deck"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"os"

	"github.com/golang/freetype"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Width  int
	Height int

	CardWidth  int
	CardHeight int

	DeckOne *deck.Deck
	DeckTwo *deck.Deck

	ShowCards   bool
	DrawOptions ebiten.DrawImageOptions

	CurrentWin int

	PlayerOneMessage string
	PlayerTwoMessage string
	PlayerOneWins    string
	PlayerTwoWins    string
	FontFace         font.Face
	FontColor        color.Color

	JustClicked      bool
	Clicking         bool
	CursorX          int
	CursorY          int
	IsCursorOnButton bool

	PlayButtonImage *ebiten.Image
}

func NewGame() *Game {
	mainDeck := deck.NewDeck()
	mainDeck.Shuffle()

	deckOne, deckTwo := mainDeck.CutInTwo()
	cardsWidth := deckOne.Cards[0].Image.Bounds().Dx()
	cardsHeight := deckOne.Cards[0].Image.Bounds().Dy()
	var drawOptions ebiten.DrawImageOptions

	// Load font
	fontFile := "assets/kongtext.ttf"
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Font
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return nil
	}

	fontFace := truetype.NewFace(f, &truetype.Options{
		Size:    50,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	fontColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	// PlayButton
	playButtonImage := ebiten.NewImage(280, 100)

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

		PlayButtonImage: playButtonImage,
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
