package main

import (
	"fmt"
	"image"
	"log"
	"testing"
)

var (
	p1 = Point{X: 100, Y: 50}
	p2 = Point{X: 30, Y: 8}
)

func TestNewPoint(t *testing.T) {
	points := []Point{
		NewPoint(1, 2),
		NewPoint(1.0, 2.0),
		NewPoint(float32(1), float32(2)),
	}

	expectation := Point{X: 1, Y: 2}

	for i := 0; i < len(points); i++ {
		point := &points[i]

		if point.X != expectation.X && point.Y != expectation.Y {
			log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, point.X, point.Y))
		}
	}
}

func TestAdd(t *testing.T) {
	p3 := p1.Add(p2.X, p2.Y)
	expectation := Point{X: 130, Y: 58}

	if p3.X != expectation.X && p3.Y != expectation.Y {
		log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, p3.X, p3.Y))
	}
}

func TestAddPoint(t *testing.T) {
	p3 := p1.AddPoint(&p2)
	expectation := Point{X: 130, Y: 58}

	if p3.X != expectation.X && p3.Y != expectation.Y {
		log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, p3.X, p3.Y))
	}
}

func TestSubtract(t *testing.T) {
	p3 := p1.Subtract(p2.X, p2.Y)
	expectation := Point{X: 70, Y: 42}

	if p3.X != expectation.X && p3.Y != expectation.Y {
		log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, p3.X, p3.Y))
	}
}

func TestSubtractPoint(t *testing.T) {
	p3 := p1.SubtractPoint(&p2)
	expectation := Point{X: 70, Y: 42}

	if p3.X != expectation.X && p3.Y != expectation.Y {
		log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, p3.X, p3.Y))
	}
}

func TestMultiply(t *testing.T) {
	p3 := p1.Multiply(p2.X, p2.Y)
	expectation := Point{X: 3000, Y: 560}

	if p3.X != expectation.X && p3.Y != expectation.Y {
		log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, p3.X, p3.Y))
	}
}

func TestMultiplyPoint(t *testing.T) {
	p3 := p1.MultiplyPoint(&p2)
	expectation := Point{X: 3000, Y: 560}

	if p3.X != expectation.X && p3.Y != expectation.Y {
		log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, p3.X, p3.Y))
	}
}

func TestFromImagePoint(t *testing.T) {
	imagePoint := image.Point{X: 1, Y: 2}
	point := FromImagePoint(imagePoint)
	expectation := Point{X: 1, Y: 2}

	if point.X != expectation.X && point.Y != expectation.Y {
		log.Fatal(fmt.Errorf("expected {X: %v, Y: %v}, but got {X: %v, Y: %v} instead", expectation.X, expectation.Y, point.X, point.Y))
	}
}
