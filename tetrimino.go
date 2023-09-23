package main

import "image"

type Tetrimino struct {
	Shape    ShapeRotations
	Position image.Point
	Rotation byte
}
