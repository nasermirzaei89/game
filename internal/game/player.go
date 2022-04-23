package game

import (
	"github.com/nasermirzaei89/td/internal/components"
)

type Player struct {
	components.Position
	components.MovementControl
	components.Character
	components.Wrapable
}
