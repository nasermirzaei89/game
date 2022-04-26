package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/game/internal/components"
	"github.com/nasermirzaei89/game/internal/engine"
)

type MovementControllable interface {
	GetPosition() *components.Position
	GetMovementControl() *components.MovementControl
}

type MovementController struct{}

var _ engine.System = new(MovementController)

func (mc *MovementController) Update(e engine.Entity) error {
	entity, ok := e.(MovementControllable)
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
