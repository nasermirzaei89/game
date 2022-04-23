package engine

import "github.com/hajimehoshi/ebiten/v2"

type Actor interface {
	Update(obj Object) error
	Draw(obj Object, screen *ebiten.Image)
}
