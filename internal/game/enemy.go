package game

import (
	"github.com/nasermirzaei89/td/internal/components"
)

type Enemy struct {
	components.Position
	components.Velocity
	components.Character
	components.Wrapable
}
