package main

import (
	"image"
	"math/rand"
)

type Tetrimino struct {
	Shape    ShapeRotations
	Position image.Point
	Rotation byte
}

func (t *Tetrimino) GetCurrentShape() *Shape {
	return &t.Shape[t.Rotation]
}

func (t *Tetrimino) Reset() {
	t.Position.X = 0
	t.Position.Y = 0
	t.Rotation = 0
	t.Shape = Shapes[rand.Intn(len(Shapes))]
}

func (t *Tetrimino) Rotate() {
	t.Rotation++

	if t.Rotation > ShapeRotationCount-1 {
		t.Rotation = 0
	}
}

func (t *Tetrimino) Left() {
	t.Position.X--
}

func (t *Tetrimino) Right() {
	t.Position.X++
}

func (t *Tetrimino) Down() {
	t.Position.Y++
}

func (t *Tetrimino) IsOccupied(x, y int) bool {
	shape := t.GetCurrentShape()
	return shape[x][y] == 1
}

type blockFn func(game *Game, x, y int) bool

func (t *Tetrimino) DoToBlocks(game *Game, fn blockFn) bool {
	ok := true

	for x := 0; x < ShapeWidth; x++ {
		for y := 0; y < ShapeHeight; y++ {
			if !t.IsOccupied(x, y) {
				continue
			}

			result := fn(game, x, y)

			if !result {
				ok = false
			}
		}
	}

	return ok
}

func (t *Tetrimino) CanMoveDown(game *Game) (isOK bool) {
	board := &game.State.Board
	position := &t.Position

	return t.DoToBlocks(game, func(game *Game, x, y int) bool {
		blockX := position.X + x
		blockY := position.Y + y

		return blockY < BoardHeight && !board.IsOccupied(blockX, blockY)
	})
}
