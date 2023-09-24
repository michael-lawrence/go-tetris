package main

import "image"

type Point struct {
	X float64
	Y float64
}

func NewPoint[T int | float32 | float64](X, Y T) Point {
	return Point{X: float64(X), Y: float64(Y)}
}

func (p *Point) Add(x, y float64) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

func (p *Point) AddPoint(p2 *Point) Point {
	return Point{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p *Point) Subtract(x, y float64) Point {
	return Point{X: p.X - x, Y: p.Y - y}
}

func (p *Point) SubtractPoint(p2 *Point) Point {
	return Point{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p *Point) Multiply(x, y float64) Point {
	return Point{X: p.X * x, Y: p.Y * y}
}

func (p *Point) MultiplyPoint(p2 *Point) Point {
	return Point{X: p.X * p2.X, Y: p.Y * p2.Y}
}

func FromImagePoint(p image.Point) Point {
	return NewPoint(p.X, p.Y)
}
