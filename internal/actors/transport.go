package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
)

type Transporter interface {
	GetPosition() *components.Position
	GetVelocity() *components.Velocity
}

type Transport struct{}

func (t *Transport) Update(obj engine.Object) error {
	entity, ok := obj.(Transporter)
	if !ok {
		return nil
	}

	position := entity.GetPosition()
	velocity := entity.GetVelocity()

	position.X += velocity.Horizontal
	position.Y += velocity.Vertical

	return nil
}

func (t *Transport) Draw(obj engine.Object, screen *ebiten.Image) {}

var _ engine.Actor = new(Transport)
