package main

import (
	"github.com/yudppp/throttle"
	"math/rand"
	"time"
)

var inputThrottler = throttle.New(time.Second / 10)

type Tetrimino struct {
	Shape    ShapeRotations
	Position Point
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
	inputThrottler.Do(func() {
		t.Position.X--
	})
}

func (t *Tetrimino) Right() {
	inputThrottler.Do(func() {
		t.Position.X++
	})
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

func (t *Tetrimino) DoToBlocks(game *Game, fn blockFn) bool {
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

	return t.DoToBlocks(game, func(game *Game, position, blockPosition Point) bool {
		return blockPosition.X > 0 && !board.Get(blockPosition.Subtract(1, 0))
	})
}

func (t *Tetrimino) CanMoveRight(game *Game) bool {
	board := &game.State.Board

	return t.DoToBlocks(game, func(game *Game, position, blockPosition Point) bool {
		return blockPosition.X < BoardWidth-1 && !board.Get(blockPosition.Add(1, 0))
	})
}

func (t *Tetrimino) CanMoveDown(game *Game) bool {
	board := &game.State.Board

	return t.DoToBlocks(game, func(game *Game, position, blockPosition Point) bool {
		return blockPosition.Y < BoardHeight-1 && !board.Get(blockPosition.Add(0, 1))
	})
}

// CanRotate @todo Implement this
func (t *Tetrimino) CanRotate(game *Game) bool {
	//board := &game.State.Board

	return t.DoToBlocks(game, func(game *Game, position, blockPosition Point) bool {
		return true
	})
}
