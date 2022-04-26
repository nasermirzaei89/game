package systems

import (
	"github.com/nasermirzaei89/game/internal/components"
	"github.com/nasermirzaei89/game/internal/engine"
)

type Transportable interface {
	GetPosition() *components.Position
	GetVelocity() *components.Velocity
}

type Transporter struct{}

var _ engine.System = new(Transporter)

func (t *Transporter) Update(obj engine.Entity) error {
	entity, ok := obj.(Transportable)
	if !ok {
		return nil
	}

	position := entity.GetPosition()
	velocity := entity.GetVelocity()

	position.X += velocity.Horizontal
	position.Y += velocity.Vertical

	return nil
}
