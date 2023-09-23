package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Game struct {
	WindowSize image.Point
	State      State
	Graphics   Graphics
}

func (game *Game) PlaceTetrimino() {
	tetrimino := &game.State.Tetrimino
	position := &tetrimino.Position
	board := &game.State.Board

	tetrimino.DoToBlocks(game, func(game *Game, x, y int) bool {
		blockX := x + position.X
		blockY := y + position.Y - 1
		board[blockX][blockY] = true
		return true
	})
}

func (game *Game) Update() error {
	tetrimino := &game.State.Tetrimino

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		tetrimino.Left()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		tetrimino.Right()
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		tetrimino.Rotate()
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		tetrimino.Down()
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		// @todo Instant drop
	}

	tetrimino.Down()

	if !tetrimino.CanDrop(game) {
		game.PlaceTetrimino()
		tetrimino.Reset()
	}

	return nil
}

func (game *Game) DrawTetrimino(screen *ebiten.Image) {
	graphic := game.Graphics.Blue
	tetrimino := &game.State.Tetrimino
	position := tetrimino.Position
	size := graphic.Bounds().Size()
	w := size.X
	h := size.Y
	x := float64(position.X * w)
	y := float64(position.Y * h)

	tetrimino.DoToBlocks(game, func(game *Game, i, j int) bool {
		blockX := x + float64(i*w)
		blockY := y + float64(j*h)

		DrawImageAt(&graphic, screen, blockX, blockY)

		return true
	})
}

func (game *Game) DrawBoard(screen *ebiten.Image) {
	board := &game.State.Board
	blank := &game.Graphics.Blank
	red := &game.Graphics.Red
	blockSize := blank.Bounds().Size()

	board.DoToBoard(game, func(game *Game, x, y int, occupied bool) bool {
		xPos := float64(x * blockSize.X)
		yPos := float64(y * blockSize.Y)

		var graphic *ebiten.Image

		if occupied {
			graphic = red
		} else {
			graphic = blank
		}

		DrawImageAt(graphic, screen, xPos, yPos)

		return true
	})
}

func DrawImageAt(source *ebiten.Image, dest *ebiten.Image, x, y float64) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(x, y)
	dest.DrawImage(source, options)
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.DrawBoard(screen)
	game.DrawTetrimino(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 200
}
