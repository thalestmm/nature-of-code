package main

import (
	"errors"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	introduction "github.com/thalestmm/nature-of-code/1_Introduction"
)

// Game implements the ebiten.Game interface and wraps around an ebiten.Game exercise
type Game struct {
	TargetFPS uint
	Exercise  ebiten.Game
}

func (g *Game) Update() error {
	if g.TargetFPS <= 0 {
		return errors.New("TargetFPS must be greater than 0")
	}

	frameDurationMilliseconds := int(1000.0 / float64(g.TargetFPS))
	time.Sleep(time.Millisecond * time.Duration(frameDurationMilliseconds))

	g.Exercise.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Exercise.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	windowWidth := 1920
	windowHeight := 1080

	game := &Game{TargetFPS: 240}

	// Exexercise 1.1:  Create a random walker that has a tendency to move down and to the right.
	ex1 := &introduction.RandomWalk{
		Location: introduction.Location{X: float64(windowWidth) / 2, Y: float64(windowHeight) / 2},
		Step:     10,
		DirectionWeights: introduction.DirectionWeights{
			Up:    1,
			Down:  1.1,
			Left:  1,
			Right: 1.1,
		},
	}
	game.Exercise = ex1

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Nature of Code")

	log.Fatal(ebiten.RunGame(game))
}
