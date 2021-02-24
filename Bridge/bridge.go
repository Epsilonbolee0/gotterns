package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
	RenderSquare(side float32)
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

func (v *VectorRenderer) RenderSquare(side float32) {
	fmt.Println("Drawing a square with side", side)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

func (r *RasterRenderer) RenderSquare(side float32) {
	fmt.Println("Drawing pixels for square with side", side)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

type Square struct {
	renderer Renderer
	side     float32
}

func NewSquare(renderer Renderer, side float32) *Square {
	return &Square{renderer, side}
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.side)
}

func (s *Square) Resize(factor float32) {
	s.side *= factor
}

func main() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	circle := NewCircle(&vector, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()

	square := NewSquare(&raster, 10)
	square.Draw()
	square.Resize(4)
	square.Draw()
}
