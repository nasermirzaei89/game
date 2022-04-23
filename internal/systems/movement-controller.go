package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
)

type MovementControllable interface {
	GetPosition() *components.Position
	GetMovementControl() *components.MovementControl
}

type MovementController struct{}

var _ engine.System = new(MovementController)

func (mc *MovementController) Update(obj engine.Entity) error {
	entity, ok := obj.(MovementControllable)
	if !ok {
		return nil
	}

	position := entity.GetPosition()
	control := entity.GetMovementControl()

	if ebiten.IsKeyPressed(control.Up) {
		position.Y -= 2
	}

	if ebiten.IsKeyPressed(control.Down) {
		position.Y += 2
	}

	if ebiten.IsKeyPressed(control.Left) {
		position.X -= 2
	}

	if ebiten.IsKeyPressed(control.Right) {
		position.X += 2
	}

	return nil
}

func (mc *MovementController) Draw(obj engine.Entity, screen *ebiten.Image) {}
