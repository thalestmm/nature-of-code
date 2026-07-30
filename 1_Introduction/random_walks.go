package introduction

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thalestmm/nature-of-code/ui"
)

type Location struct {
	X float64
	Y float64
}

type RandomWalk struct {
	Location Location
	Path     []Location
	Step     int
}

func (rw *RandomWalk) Update() error {
	if rw.Step == 0 {
		rw.Step = 1
	}

	// Append the current location to the path
	rw.Path = append(rw.Path, rw.Location)

	flip1 := rand.Float64() > 0.5
	flip2 := rand.Float64() > 0.5

	// Next step logic
	if flip1 && flip2 {
		rw.Location.Y += float64(rw.Step)
	}
	if flip1 && !flip2 {
		rw.Location.X += float64(rw.Step)
	}
	if !flip1 && flip2 {
		rw.Location.X -= float64(rw.Step)
	}
	if !flip1 && !flip2 {
		rw.Location.Y -= float64(rw.Step)
	}

	return nil
}

func (rw *RandomWalk) Draw(screen *ebiten.Image) {
	backgroundColor := color.RGBA{254, 249, 255, 255}
	screen.Fill(backgroundColor)

	// Draw all previous steps
	for _, loc := range rw.Path {
		square := ui.Square{Width: 10, Color: color.RGBA{0, 0, 0, 255}, PosX: loc.X, PosY: loc.Y}
		screen.DrawImage(square.Image(), square.Opts)
	}

	// Draw the current location
	square := ui.Square{Width: 10, Color: color.RGBA{255, 0, 0, 255}, PosX: rw.Location.X, PosY: rw.Location.Y}
	screen.DrawImage(square.Image(), square.Opts)
}

func (rw *RandomWalk) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
