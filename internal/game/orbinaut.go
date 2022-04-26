package game

import (
	"github.com/nasermirzaei89/game/internal/components"
)

type Orbinaut struct {
	components.Position
	components.Velocity
	components.Animation
	components.Wrapable
}
