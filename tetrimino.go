package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Tetrimino struct {
	Position Point
	ShapeNum int
	Shape    Shape
}

func (t *Tetrimino) GetCurrentShape() *Shape {
	return &t.Shape
}

func (t *Tetrimino) GetCurrentGraphic() *ebiten.Image {
	return &Colors[t.ShapeNum+1]
}

func (t *Tetrimino) Reset(game *Game) {
	t.ShapeNum = rand.Intn(len(Shapes))
	t.Shape = Shapes[t.ShapeNum]
	t.Position.X = 0
	t.Position.Y = 0
}

func (t *Tetrimino) Rotate() {
	t.Shape.Rotate()
}

func (t *Tetrimino) Left() {
	t.Position.X--
}

func (t *Tetrimino) Right() {
	t.Position.X++
}

func (t *Tetrimino) Down() {
	t.Position.Y += 0.1
}

func (t *Tetrimino) Drop(game *Game) {
	for t.CanMoveDown(game) {
		t.Down()
	}
}

func (t *Tetrimino) Get(position Point) byte {
	shape := t.GetCurrentShape()
	return shape[int(position.X)][int(position.Y)]
}

func (t *Tetrimino) Set(position Point, value byte) byte {
	shape := t.GetCurrentShape()
	shape[int(position.X)][int(position.Y)] = value
	return value
}

func (t *Tetrimino) IsOccupied(position Point) bool {
	return t.Get(position) == 1
}

type blockFn func(game *Game, position, blockPosition Point) bool

func (t *Tetrimino) Do(game *Game, fn blockFn) bool {
	ok := true
	position := &t.Position

	for i := 0; i < ShapeWidth; i++ {
		for j := 0; j < ShapeHeight; j++ {
			pos := NewPoint(i, j)
			blockPos := pos.AddPoint(position)

			if !t.IsOccupied(pos) {
				continue
			}

			result := fn(game, pos, blockPos)

			if !result {
				ok = false
			}
		}
	}

	return ok
}

func (t *Tetrimino) CanMoveLeft(game *Game) bool {
	board := &game.State.Board

	return t.Do(game, func(game *Game, position, blockPosition Point) bool {
		return blockPosition.X > 0 && !board.IsOccupied(blockPosition.Subtract(1, 0))
	})
}

func (t *Tetrimino) CanMoveRight(game *Game) bool {
	board := &game.State.Board

	return t.Do(game, func(game *Game, position, blockPosition Point) bool {
		return blockPosition.X < BoardWidth-1 && !board.IsOccupied(blockPosition.Add(1, 0))
	})
}

func (t *Tetrimino) CanMoveDown(game *Game) bool {
	board := &game.State.Board

	return t.Do(game, func(game *Game, position, blockPosition Point) bool {
		return blockPosition.Y < BoardHeight-1 && !board.IsOccupied(blockPosition.Add(0, 1))
	})
}

// CanRotate @todo Implement this
func (t *Tetrimino) CanRotate(game *Game) bool {
	//board := &game.State.Board

	return t.Do(game, func(game *Game, position, blockPosition Point) bool {
		return true
	})
}
