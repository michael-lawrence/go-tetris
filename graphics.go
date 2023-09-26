package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
	"os"
)

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
