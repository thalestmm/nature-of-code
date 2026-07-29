package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thalestmm/nature-of-code/ui"
)

// Game implements the ebiten.Game interface
type Game struct {
	Entities []ui.Entity
}

func (g *Game) Update() error {
	time.Sleep(time.Millisecond * 50)

	for _, entity := range g.Entities {
		entity.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	backgroundColor := color.RGBA{241, 143, 1, 255}
	screen.Fill(backgroundColor)

	for _, entity := range g.Entities {
		screen.DrawImage(entity.Image(), entity.Options())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	circle := ui.Circle{Radius: 20, Color: color.RGBA{0, 0, 0, 255}, PosX: 200.0, PosY: 200.0}

	game := &Game{Entities: []ui.Entity{&circle}}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Nature of Code")

	log.Fatal(ebiten.RunGame(game))
}
