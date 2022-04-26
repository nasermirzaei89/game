package game

import (
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
)

type Orbinaut struct {
	components2.Position
	components2.Velocity
	components2.Animation
	components2.Wrapable
}
