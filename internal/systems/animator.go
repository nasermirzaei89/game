package systems

import (
	"math"

	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
)

type Animation interface {
	GetPosition() *components.Position
	GetSprite() *components.Sprite
	GetAnimation() *components.Animation
}

type Animator struct{}

var _ engine.System = new(Animator)

func (a *Animator) Update(e engine.Entity) error {
	entity, ok := e.(Animation)
	if !ok {
		return nil
	}

	animation := entity.GetAnimation()

	animation.FrameTime += animation.FrameRate / 60.0
	animation.FrameTime = math.Mod(animation.FrameTime, float64(len(animation.Frames)))

	return nil
}
