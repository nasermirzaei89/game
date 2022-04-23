package game

import (
	"github.com/nasermirzaei89/td/internal/components"
)

type Orbinaut struct {
	components.Position
	components.Velocity
	components.Animation
	components.Wrapable
}
