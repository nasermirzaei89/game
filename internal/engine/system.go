package engine

import "github.com/hajimehoshi/ebiten/v2"

type System interface {
	Update(obj Entity) error
	Draw(obj Entity, screen *ebiten.Image)
}
