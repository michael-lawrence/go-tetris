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

func (t *Tetrimino) Drop(game *Game) {
	for t.CanMoveDown(game) {
		t.Down()
	}
}

func (t *Tetrimino) IsOccupied(x, y int) bool {
	shape := t.GetCurrentShape()
	return shape[x][y] == 1
}

type blockFn func(game *Game, x, y, blockX, blockY int) bool

func (t *Tetrimino) DoToBlocks(game *Game, fn blockFn) bool {
	ok := true
	position := &t.Position

	for x := 0; x < ShapeWidth; x++ {
		for y := 0; y < ShapeHeight; y++ {
			if !t.IsOccupied(x, y) {
				continue
			}

			blockX := x + position.X
			blockY := y + position.Y

			result := fn(game, x, y, blockX, blockY)

			if !result {
				ok = false
			}
		}
	}

	return ok
}

func (t *Tetrimino) CanMoveLeft(game *Game) bool {
	board := &game.State.Board

	return t.DoToBlocks(game, func(game *Game, x, y, blockX, blockY int) bool {
		return blockX > 0 && !board.IsOccupied(blockX-1, blockY)
	})
}

func (t *Tetrimino) CanMoveRight(game *Game) bool {
	board := &game.State.Board

	return t.DoToBlocks(game, func(game *Game, x, y, blockX, blockY int) bool {
		return blockX < BoardWidth-1 && !board.IsOccupied(blockX+1, blockY)
	})
}

func (t *Tetrimino) CanMoveDown(game *Game) bool {
	board := &game.State.Board

	return t.DoToBlocks(game, func(game *Game, x, y, blockX, blockY int) bool {
		return blockY < BoardHeight-1 && !board.IsOccupied(blockX, blockY+1)
	})
}

// CanRotate @todo Implement this
func (t *Tetrimino) CanRotate(game *Game) bool {
	//board := &game.State.Board

	return t.DoToBlocks(game, func(game *Game, x, y, blockX, blockY int) bool {
		return true
	})
}
