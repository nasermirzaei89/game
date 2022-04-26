package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
	engine2 "github.com/nasermirzaei89/game/internal/game1/engine"
)

type MovementControllable interface {
	GetPosition() *components2.Position
	GetMovementControl() *components2.MovementControl
}

type MovementController struct{}

var _ engine2.System = new(MovementController)

func (mc *MovementController) Update(e engine2.Entity) error {
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
