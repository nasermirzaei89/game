package engine

import "github.com/hajimehoshi/ebiten/v2"

type System interface{}

type Registrable interface {
	Register(p *Project)
}

type Updatable interface {
	Update(entity Entity) error
}

type Drawable interface {
	Draw(entity Entity, screen *ebiten.Image)
}
