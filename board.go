package main

const (
	BoardWidth  = 10
	BoardHeight = 20
)

type Board [BoardWidth][BoardHeight]bool

type boardFn func(game *Game, x, y int, occupied bool) bool

func (b *Board) IsOccupied(x, y int) bool {
	return b[x][y]
}

func (b *Board) DoToBoard(game *Game, fn boardFn) bool {
	ok := true

	for x := 0; x < BoardWidth; x++ {
		for y := 0; y < BoardHeight; y++ {
			ok = ok && fn(game, x, y, b.IsOccupied(x, y))
		}
	}

	return ok
}
