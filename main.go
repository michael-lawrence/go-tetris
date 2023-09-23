package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"math/rand"
)

var (
	shapes = []ShapeRotations{I, J, L, O, S, T, Z}
	game   = &Game{
		WindowSize: image.Point{X: 400, Y: 800},
		State: State{
			Tetrimino: Tetrimino{
				Shape:    shapes[0],
				Position: image.Point{X: 0, Y: 0},
				Rotation: 0,
			},
			Board: [10][20]byte{},
		},
		Graphics: Graphics{
			Blank:  *LoadImage("resources/images/blank.png"),
			Blue:   *LoadImage("resources/images/blue.png"),
			Green:  *LoadImage("resources/images/green.png"),
			Orange: *LoadImage("resources/images/orange.png"),
			Purple: *LoadImage("resources/images/purple.png"),
			Red:    *LoadImage("resources/images/red.png"),
		},
	}
)

func (game *Game) Update() error {
	tetrimino := &game.State.Tetrimino
	position := &tetrimino.Position

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		position.X--
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		position.X++
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		tetrimino.Rotation++

		if tetrimino.Rotation > 3 {
			tetrimino.Rotation = 0
		}
	}

	position.Y++

	if position.Y >= len(game.State.Board[0]) {
		position.X = 0
		position.Y = 0
		tetrimino.Rotation = 0
		tetrimino.Shape = shapes[rand.Intn(len(shapes))]
	}

	return nil
}

func (game *Game) DrawTetrimino(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	graphic := game.Graphics.Blue
	tetrimino := &game.State.Tetrimino
	position := tetrimino.Position
	shape := tetrimino.Shape[tetrimino.Rotation]
	size := graphic.Bounds().Size()
	w := size.X
	h := size.Y
	x := float64(position.X * w)
	y := float64(position.Y * h)

	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == 1 {
				blockX := x + float64(i*w)
				blockY := y + float64(j*h)

				options.GeoM.Reset()
				options.GeoM.Translate(blockX, blockY)
				screen.DrawImage(&graphic, options)
			}
		}
	}
}

func (game *Game) DrawBoard(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	blockSize := game.Graphics.Blank.Bounds().Size()

	for x := 0; x < len(game.State.Board); x++ {
		xPos := float64(x * blockSize.X)

		for y := 0; y < len(game.State.Board[x]); y++ {
			yPos := float64(y * blockSize.Y)

			options.GeoM.Reset()
			options.GeoM.Translate(xPos, yPos)

			if game.State.Board[x][y] == 0 {
				screen.DrawImage(&game.Graphics.Blank, options)
			} else {
				screen.DrawImage(&game.Graphics.Red, options)
			}
		}
	}
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.DrawBoard(screen)
	game.DrawTetrimino(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 200
}

func main() {
	ebiten.SetWindowSize(game.WindowSize.X, game.WindowSize.Y)
	ebiten.SetWindowTitle("Tetris")
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetTPS(10)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
