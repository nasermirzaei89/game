package game

import (
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
)

type BatBrain struct {
	components2.Position
	components2.MovementControl
	components2.Animation
	components2.Wrapable
}
