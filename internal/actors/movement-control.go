package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
)

type MovementController interface {
	GetPosition() *components.Position
	GetMovementControl() *components.MovementControl
}

type MovementControl struct{}

func (mc *MovementControl) Update(obj engine.Object) error {
	entity, ok := obj.(MovementController)
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

func (mc *MovementControl) Draw(obj engine.Object, screen *ebiten.Image) {}

var _ engine.Actor = new(MovementControl)
