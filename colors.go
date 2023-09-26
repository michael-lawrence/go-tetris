package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	Colors = [8]ebiten.Image{
		*loadImage("resources/images/blank.png"),
		*loadImage("resources/images/red.png"),
		*loadImage("resources/images/orange.png"),
		*loadImage("resources/images/yellow.png"),
		*loadImage("resources/images/green.png"),
		*loadImage("resources/images/blue.png"),
		*loadImage("resources/images/purple.png"),
		*loadImage("resources/images/teal.png"),
	}
)
