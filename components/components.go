package components

type Position struct {
	X, Y int
}

type Texture struct {
	Name  string
	Order int // rendering order
}

type MoveInput struct {
	Dx, Dy int
}
