package introduction

import (
	"errors"
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thalestmm/nature-of-code/ui"
)

type Location struct {
	X float64
	Y float64
}

type DirectionWeights struct {
	Up    float64
	Down  float64
	Left  float64
	Right float64
}

type RandomWalk struct {
	Location         Location
	Path             []Location
	Step             int
	DirectionWeights DirectionWeights
}

func (rw *RandomWalk) Update() error {
	if rw.Step == 0 {
		rw.Step = 1
	}

	totalWeights := rw.DirectionWeights.Up + rw.DirectionWeights.Down + rw.DirectionWeights.Left + rw.DirectionWeights.Right

	if totalWeights == 0 {
		return errors.New("total weights cannot be zero")
	}

	// Append the current location to the path
	rw.Path = append(rw.Path, rw.Location)

	ySlider := rand.Float64()
	xSlider := rand.Float64()

	upDownRatio := rw.DirectionWeights.Up / rw.DirectionWeights.Down
	leftRightRatio := rw.DirectionWeights.Left / rw.DirectionWeights.Right

	upCutoff := 0.5 * upDownRatio
	leftCutoff := 0.5 * leftRightRatio

	up := ySlider < upCutoff
	left := xSlider < leftCutoff

	// TODO: Add logic to take a single step per iteration
	if up {
		rw.Location.Y -= float64(rw.Step)
	} else {
		rw.Location.Y += float64(rw.Step)
	}
	if left {
		rw.Location.X -= float64(rw.Step)
	} else {
		rw.Location.X += float64(rw.Step)
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
