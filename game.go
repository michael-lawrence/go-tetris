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
	board := &game.State.Board

	tetrimino.DoToBlocks(game, func(game *Game, position, blockPosition Point) bool {
		return board.Set(blockPosition, true)
	})
}

func (game *Game) HandleKeyboard() {
	tetrimino := &game.State.Tetrimino

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && tetrimino.CanMoveLeft(game) {
		tetrimino.Left()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && tetrimino.CanMoveRight(game) {
		tetrimino.Right()
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) && tetrimino.CanRotate(game) {
		tetrimino.Rotate()
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && tetrimino.CanMoveDown(game) {
		tetrimino.Down()
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		tetrimino.Drop(game)
	}
}

// GameOver @todo Make game over screen
func (game *Game) GameOver() {
	game.State.Board = Board{}
}

func (game *Game) Update() error {
	tetrimino := &game.State.Tetrimino

	if !tetrimino.CanMoveDown(game) {
		game.GameOver()
	}

	tetrimino.Down()
	game.HandleKeyboard()

	if !tetrimino.CanMoveDown(game) {
		game.PlaceTetrimino()
		tetrimino.Reset()
	}

	return nil
}

func (game *Game) DrawTetrimino(screen *ebiten.Image) {
	graphic := game.Graphics.Blue
	tetrimino := &game.State.Tetrimino
	blockSize := FromImagePoint(graphic.Bounds().Size())

	tetrimino.DoToBlocks(game, func(game *Game, position, blockPosition Point) bool {
		currentBlockPosition := blockPosition.MultiplyPoint(&blockSize)
		DrawImageAt(&graphic, screen, currentBlockPosition)

		return true
	})
}

func (game *Game) DrawBoard(screen *ebiten.Image) {
	board := &game.State.Board
	blank := &game.Graphics.Blank
	red := &game.Graphics.Red
	blockSize := FromImagePoint(blank.Bounds().Size())

	board.DoToBoard(game, func(game *Game, position Point, occupied bool) bool {
		currentBlockPosition := blockSize.MultiplyPoint(&position)

		var graphic *ebiten.Image

		if occupied {
			graphic = red
		} else {
			graphic = blank
		}

		DrawImageAt(graphic, screen, currentBlockPosition)

		return true
	})
}

func DrawImageAt(source *ebiten.Image, dest *ebiten.Image, position Point) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(position.X, position.Y)
	dest.DrawImage(source, options)
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.DrawBoard(screen)
	game.DrawTetrimino(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 200
}
