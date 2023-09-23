package main

import "image"

type Game struct {
	WindowSize image.Point
	State      State
	Graphics   Graphics
}
