package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
	engine2 "github.com/nasermirzaei89/game/internal/game1/engine"
)

type Drawable interface {
	GetPosition() *components2.Position
	GetSprite() *components2.Sprite
}

type Drawer struct{}

var _ engine2.System = new(Drawer)

func (d *Drawer) Draw(e engine2.Entity, screen *ebiten.Image) {
	entity, ok := e.(Drawable)
	if !ok {
		return
	}

	position := entity.GetPosition()
	sprite := entity.GetSprite()

	ops := &ebiten.DrawImageOptions{}

	ops.GeoM.Translate(float64(-sprite.OffsetX), float64(-sprite.OffsetY))

	ops.GeoM.Translate(position.X, position.Y)

	screen.DrawImage(sprite.Image, ops)
}
