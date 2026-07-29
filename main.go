package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements the ebiten.Game interface
type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	backgroundColor := color.RGBA{241, 143, 1, 255}
	screen.Fill(backgroundColor)

	// Render a simple square
	entityColor := color.RGBA{0, 0, 0, 255}
	square := ebiten.NewImage(10, 10)
	square.Fill(entityColor)

	op := &ebiten.DrawImageOptions{}
	rx := rand.Float64()
	ry := rand.Float64()
	op.GeoM.Translate(rx*800.0, ry*600.0)

	screen.DrawImage(square, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Nature of Code")

	log.Fatal(ebiten.RunGame(game))
}
