package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yudppp/throttle"
	"image"
	"time"
)

var inputThrottler = throttle.New(time.Second / 10)

type Game struct {
	WindowSize image.Point
	State      State
	Keyboard   Keyboard
}

func (game *Game) PlaceTetrimino() {
	tetrimino := &game.State.Tetrimino
	board := &game.State.Board

	tetrimino.Do(game, func(game *Game, position, blockPosition Point) bool {
		return board.Set(blockPosition, tetrimino.ShapeNum+1) > 0
	})
}

// GameOver @todo Make game over screen
func (game *Game) GameOver() {
	game.State.Board = Board{}
}

func (game *Game) Update() error {
	tetrimino := &game.State.Tetrimino
	keyboard := &game.Keyboard

	if !tetrimino.CanMoveDown(game) {
		game.GameOver()
	}

	tetrimino.Down()
	keyboard.Process(game)

	if !tetrimino.CanMoveDown(game) {
		game.PlaceTetrimino()
		tetrimino.Reset(game)
	}

	return nil
}

func (game *Game) DrawTetrimino(screen *ebiten.Image) {
	tetrimino := &game.State.Tetrimino
	graphic := tetrimino.GetCurrentGraphic()
	blockSize := FromImagePoint(graphic.Bounds().Size())

	tetrimino.Do(game, func(game *Game, position, blockPosition Point) bool {
		currentBlockPosition := blockPosition.MultiplyPoint(&blockSize)
		DrawImageAt(graphic, screen, currentBlockPosition)

		return true
	})
}

func (game *Game) DrawBoard(screen *ebiten.Image) {
	board := &game.State.Board

	board.Do(game, func(game *Game, position Point, occupied bool) bool {
		graphic := &Colors[board.Get(position)]
		blockSize := FromImagePoint(graphic.Bounds().Size())
		currentBlockPosition := blockSize.MultiplyPoint(&position)

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
