package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Entity interface {
	Image() *ebiten.Image
	Options() *ebiten.DrawImageOptions
}

type Square struct {
	Width int
	Color color.Color
	Opts  *ebiten.DrawImageOptions
	PosX  float64
	PosY  float64
}

func (s *Square) Image() *ebiten.Image {
	square := ebiten.NewImage(s.Width, s.Width)
	square.Fill(s.Color)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.PosX-float64(s.Width/2), s.PosY-float64(s.Width/2))
	s.Opts = op

	return square
}

func (s *Square) Options() *ebiten.DrawImageOptions {
	return s.Opts
}

type Circle struct {
	Radius int
	Color  color.Color
	Opts   *ebiten.DrawImageOptions
	PosX   float64
	PosY   float64
}

func (c *Circle) Image() *ebiten.Image {
	circle := ebiten.NewImage(c.Radius*2, c.Radius*2)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.PosX-float64(c.Radius), c.PosY-float64(c.Radius))

	vector.FillCircle(circle, float32(c.Radius), float32(c.Radius), float32(c.Radius), c.Color, true)

	c.Opts = op

	return circle
}

func (c *Circle) Options() *ebiten.DrawImageOptions {
	return c.Opts
}
