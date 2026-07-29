package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/thalestmm/nature-of-code/ui"
)

// Game implements the ebiten.Game interface
type Game struct {
}

func (g *Game) Update() error {
	time.Sleep(time.Millisecond * 16)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	backgroundColor := color.RGBA{241, 143, 1, 255}
	screen.Fill(backgroundColor)

	// Render a simple square
	square := ui.Square{Width: 10, Color: color.RGBA{0, 0, 0, 255}, PosX: 100.0, PosY: 100.0}
	sqImage := square.Image()
	opSquare := square.Options()

	circle := ui.Circle{Radius: 20, Color: color.RGBA{0, 0, 0, 255}, PosX: 200.0, PosY: 200.0}
	circleImage := circle.Image()
	opCircle := circle.Options()

	screen.DrawImage(sqImage, opSquare)
	screen.DrawImage(circleImage, opCircle)

	vector.FillCircle(screen, 100, 200, 20, color.RGBA{255, 255, 255, 255}, true)
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
