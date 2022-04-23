package game

import (
	"github.com/nasermirzaei89/td/internal/components"
)

type BatBrain struct {
	components.Position
	components.MovementControl
	components.Animation
	components.Wrapable
}
