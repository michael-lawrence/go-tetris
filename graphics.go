package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
	"os"
)

type Graphics struct {
	Blank  ebiten.Image
	Blue   ebiten.Image
	Green  ebiten.Image
	Orange ebiten.Image
	Purple ebiten.Image
	Red    ebiten.Image
}

func NewGraphics() *Graphics {
	return &Graphics{
		Blank:  *loadImage("resources/images/blank.png"),
		Blue:   *loadImage("resources/images/blue.png"),
		Green:  *loadImage("resources/images/green.png"),
		Orange: *loadImage("resources/images/orange.png"),
		Purple: *loadImage("resources/images/purple.png"),
		Red:    *loadImage("resources/images/red.png"),
	}
}

func loadImage(path string) (img *ebiten.Image) {
	imgFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	imgReader, _, err := image.Decode(imgFile)

	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(imgReader)
}
