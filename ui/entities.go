package ui

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Entity interface {
	Image() *ebiten.Image
	Update()
	Options() *ebiten.DrawImageOptions
}

type Square struct {
	Width int
	Color color.Color
	PosX  float64
	PosY  float64
	Opts  *ebiten.DrawImageOptions
	image *ebiten.Image
}

func (s *Square) Image() *ebiten.Image {
	square := ebiten.NewImage(s.Width, s.Width)
	square.Fill(s.Color)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.PosX-float64(s.Width/2), s.PosY-float64(s.Width/2))
	s.Opts = op

	return square
}

func (s *Square) Update() {
}

func (s *Square) Options() *ebiten.DrawImageOptions {
	return s.Opts
}

type Circle struct {
	Radius int
	Color  color.Color
	PosX   float64
	PosY   float64
	Opts   *ebiten.DrawImageOptions
}

func (c *Circle) Image() *ebiten.Image {
	circle := ebiten.NewImage(c.Radius*2, c.Radius*2)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.PosX-float64(c.Radius), c.PosY-float64(c.Radius))

	vector.FillCircle(circle, float32(c.Radius), float32(c.Radius), float32(c.Radius), c.Color, true)

	c.Opts = op

	return circle
}

func (c *Circle) Update() {
	c.PosX = 800.0 * rand.Float64()
	c.PosY = 600.0 * rand.Float64()
}

func (c *Circle) Options() *ebiten.DrawImageOptions {
	return c.Opts
}

type Emitter struct {
	Entity  Entity
	Color   color.Color
	Radius  int
	Density int
	PosX    float64
	PosY    float64
	Opts    *ebiten.DrawImageOptions
}

// TODO: Implement
func (e *Emitter) Image() *ebiten.Image {
	emitter := ebiten.NewImage(e.Radius*2, e.Radius*2)

	// Spawn n (density) particles in random directions
	// that fade away over time up to a given distance (radius)
	centerX := e.PosX - float64(e.Radius)
	centerY := e.PosY - float64(e.Radius)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(centerX, centerY)
	e.Opts = op

	return emitter
}

func (e *Emitter) Options() *ebiten.DrawImageOptions {
	return e.Opts
}
