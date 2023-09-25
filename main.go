package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

var (
	game = &Game{
		WindowSize: image.Point{X: 400, Y: 800},
		State: State{
			Tetrimino: Tetrimino{
				Shape:    Shapes[0],
				Position: Point{X: 0, Y: 0},
			},
			Board: Board{},
		},
		Graphics: *NewGraphics(),
		Keyboard: Keyboard{
			Left:   ebiten.KeyLeft,
			Right:  ebiten.KeyRight,
			Down:   ebiten.KeyDown,
			Rotate: ebiten.KeyUp,
			Drop:   ebiten.KeySpace,
		},
	}
)

func main() {
	ebiten.SetWindowSize(game.WindowSize.X, game.WindowSize.Y)
	ebiten.SetWindowTitle("Tetris")
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetTPS(60)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
