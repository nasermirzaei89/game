package components

type Position struct {
	X, Y float64
}

func (p *Position) GetPosition() *Position {
	return p
}
