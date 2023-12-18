package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawDeckOne(screen)
	g.DrawDeckTwo(screen)

	// Draw messages
	text.Draw(screen, g.PlayerTwoMessage, g.FontFace, g.Width*9/10, g.Height*1/10, g.FontColor)
	text.Draw(screen, g.PlayerOneMessage, g.FontFace, g.Width*9/10, g.Height*9/10, g.FontColor)

	// Draw win badge
	text.Draw(screen, g.PlayerTwoWins, g.FontFace, g.Width*2/3, g.Height*1/3, g.FontColor)
	text.Draw(screen, g.PlayerOneWins, g.FontFace, g.Width*2/3, g.Height*2/3, g.FontColor)

	g.DrawPlayButton(screen)

	return
}

func (g *Game) DrawDeckOne(screen *ebiten.Image) {
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

}

func (g *Game) DrawDeckTwo(screen *ebiten.Image) {
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
}

func (g *Game) DrawPlayButton(screen *ebiten.Image) {
	buttonColor := color.RGBA{R: 150, B: 150, A: 200}

	vector.DrawFilledRect(
		screen,
		float32(g.PlayButtonX)-5,
		float32(g.PlayButtonY)-5,
		float32(g.PlayButtonImage.Bounds().Dx())+10,
		float32(g.PlayButtonImage.Bounds().Dy())+10,
		color.NRGBA{R: 255, G: 255, B: 255, A: uint8(g.PlayButtonHoverAlpha)},
		true,
	)

	if g.Clicking {
		buttonColor = color.RGBA{R: 150, B: 150, A: 150}
	} else {
		buttonColor = color.RGBA{R: 200, B: 200, A: 200}
	}

	g.PlayButtonImage.Fill(buttonColor)

	g.DrawOptions.GeoM.Reset()
	g.DrawOptions.GeoM.Translate(g.PlayButtonX, g.PlayButtonY)

	screen.DrawImage(g.PlayButtonImage, &g.DrawOptions)
	text.Draw(
		screen,
		"Play",
		g.FontFace,
		int(g.PlayButtonX*1.04),
		int(g.PlayButtonY*1.23),
		g.FontColor,
	)
}
