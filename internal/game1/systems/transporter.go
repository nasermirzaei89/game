package systems

import (
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
	engine2 "github.com/nasermirzaei89/game/internal/game1/engine"
)

type Transportable interface {
	GetPosition() *components2.Position
	GetVelocity() *components2.Velocity
}

type Transporter struct{}

var _ engine2.System = new(Transporter)

func (t *Transporter) Update(obj engine2.Entity) error {
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
