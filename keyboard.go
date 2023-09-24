package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct {
	Left   ebiten.Key
	Right  ebiten.Key
	Down   ebiten.Key
	Rotate ebiten.Key
	Drop   ebiten.Key
}

func (k *Keyboard) Process(game *Game) {
	tetrimino := &game.State.Tetrimino

	if ebiten.IsKeyPressed(k.Left) && tetrimino.CanMoveLeft(game) {
		inputThrottler.Do(func() {
			tetrimino.Left()
		})
	} else if ebiten.IsKeyPressed(k.Right) && tetrimino.CanMoveRight(game) {
		inputThrottler.Do(func() {
			tetrimino.Right()
		})
	} else if inpututil.IsKeyJustPressed(k.Rotate) && tetrimino.CanRotate(game) {
		tetrimino.Rotate()
	} else if ebiten.IsKeyPressed(k.Down) && tetrimino.CanMoveDown(game) {
		tetrimino.Down()
	} else if inpututil.IsKeyJustPressed(k.Drop) {
		tetrimino.Drop(game)
	}
}
