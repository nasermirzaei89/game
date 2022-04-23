package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
)

type Drawable interface {
	GetPosition() *components.Position
	GetCharacter() *components.Character
}

type Draw struct{}

func (d *Draw) Update(obj engine.Object) error {
	return nil
}

func (d *Draw) Draw(obj engine.Object, screen *ebiten.Image) {
	entity, ok := obj.(Drawable)
	if !ok {
		return
	}

	position := entity.GetPosition()
	character := entity.GetCharacter()

	ebitenutil.DebugPrintAt(screen, string(*character), int(position.X), int(position.Y))
}

var _ engine.Actor = new(Draw)
