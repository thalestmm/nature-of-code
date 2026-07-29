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
	Entities  []ui.Entity
	TargetFPS uint
}

func (g *Game) Update() error {
	frameDurationMilliseconds := int(1000.0 / float64(g.TargetFPS))
	time.Sleep(time.Millisecond * time.Duration(frameDurationMilliseconds))

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
	circle := ui.Circle{Radius: 20, Color: color.RGBA{0, 0, 0, 255}, PosX: 400.0, PosY: 300.0}
	square1 := ui.Square{Width: 10, Color: color.RGBA{255, 255, 255, 255}, PosX: 5, PosY: 5}
	square2 := ui.Square{Width: 10, Color: color.RGBA{255, 255, 255, 255}, PosX: 15, PosY: 15}

	game := &Game{Entities: []ui.Entity{&circle, &square1, &square2}, TargetFPS: 120}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Nature of Code")

	log.Fatal(ebiten.RunGame(game))
}
