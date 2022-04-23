package components

import "github.com/hajimehoshi/ebiten/v2"

type MovementControl struct {
	Up    ebiten.Key
	Down  ebiten.Key
	Left  ebiten.Key
	Right ebiten.Key
}

func (mc *MovementControl) GetMovementControl() *MovementControl {
	return mc
}
