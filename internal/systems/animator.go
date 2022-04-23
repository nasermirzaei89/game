package systems

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
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

func (a *Animator) Update(obj engine.Entity) error {
	entity, ok := obj.(Animation)
	if !ok {
		return nil
	}

	animation := entity.GetAnimation()

	animation.FrameTime += animation.FrameRate / 60.0
	animation.FrameTime = math.Mod(animation.FrameTime, float64(len(animation.Frames)))

	return nil
}

func (a *Animator) Draw(obj engine.Entity, screen *ebiten.Image) {
	entity, ok := obj.(Drawable)
	if !ok {
		return
	}

	position := entity.GetPosition()
	sprite := entity.GetSprite()

	ops := &ebiten.DrawImageOptions{}

	ops.GeoM.Translate(float64(-sprite.OffsetX), float64(-sprite.OffsetY))

	ops.GeoM.Translate(position.X, position.Y)

	screen.DrawImage(sprite.Image(), ops)
}
