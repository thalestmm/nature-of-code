package introduction

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thalestmm/nature-of-code/ui"
)

type RandomWalk struct {
	X    float64
	Y    float64
	Step int
}

func (rw *RandomWalk) Update() error {
	if rw.Step == 0 {
		rw.Step = 1
	}

	flip1 := rand.Float64() > 0.5
	flip2 := rand.Float64() > 0.5

	if flip1 && flip2 {
		rw.Y += float64(rw.Step)
	}
	if flip1 && !flip2 {
		rw.X += float64(rw.Step)
	}
	if !flip1 && flip2 {
		rw.X -= float64(rw.Step)
	}
	if !flip1 && !flip2 {
		rw.Y -= float64(rw.Step)
	}

	return nil
}

func (rw *RandomWalk) Draw(screen *ebiten.Image) {
	backgroundColor := color.RGBA{254, 249, 255, 255}
	screen.Fill(backgroundColor)

	square := ui.Square{Width: 10, Color: color.RGBA{0, 0, 0, 255}, PosX: rw.X, PosY: rw.Y}
	screen.DrawImage(square.Image(), square.Opts)
}

func (rw *RandomWalk) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
