package systems

import (
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
	engine2 "github.com/nasermirzaei89/game/internal/game1/engine"
	"math"
)

type Animation interface {
	GetPosition() *components2.Position
	GetSprite() *components2.Sprite
	GetAnimation() *components2.Animation
}

type Animator struct{}

var _ engine2.System = new(Animator)

func (a *Animator) Update(e engine2.Entity) error {
	entity, ok := e.(Animation)
	if !ok {
		return nil
	}

	animation := entity.GetAnimation()

	animation.FrameTime += animation.FrameRate / 60.0
	animation.FrameTime = math.Mod(animation.FrameTime, float64(len(animation.Frames)))

	return nil
}
