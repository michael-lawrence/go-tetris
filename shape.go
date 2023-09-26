package main

const (
	ShapeWidth  = 4
	ShapeHeight = 4
)

type Shape [ShapeWidth][ShapeHeight]byte

func (s *Shape) Rotate() *Shape {
	shape := s
	n := len(shape)

	// transpose the s
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp := shape[i][j]
			shape[i][j] = shape[j][i]
			shape[j][i] = temp
		}
	}

	// reverse each column
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			temp := shape[n-j-1][i]
			shape[n-j-1][i] = shape[j][i]
			shape[j][i] = temp
		}
	}

	return shape
}

var (
	I = Shape{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	J = Shape{
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	L = Shape{
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	O = Shape{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	S = Shape{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	T = Shape{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 1, 1},
		{0, 0, 0, 0},
	}
	Z = Shape{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 0, 0},
	}

	Shapes = []Shape{L, O, S, T, J, I, Z}
)
