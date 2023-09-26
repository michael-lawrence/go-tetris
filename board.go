package main

const (
	BoardWidth  = 10
	BoardHeight = 20
)

type Board [BoardWidth][BoardHeight]int

type boardFn func(game *Game, position Point, occupied bool) bool

func (b *Board) Set(position Point, value int) int {
	b[int(position.X)][int(position.Y)] = value
	return value
}

func (b *Board) Get(position Point) int {
	return b[int(position.X)][int(position.Y)]
}

func (b *Board) IsOccupied(position Point) bool {
	return b.Get(position) > 0
}

func (b *Board) Do(game *Game, fn boardFn) bool {
	ok := true

	for x := 0; x < BoardWidth; x++ {
		for y := 0; y < BoardHeight; y++ {
			position := NewPoint(x, y)
			ok = ok && fn(game, position, b.IsOccupied(position))
		}
	}

	return ok
}
