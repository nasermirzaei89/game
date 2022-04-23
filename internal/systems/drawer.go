package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
)

type Drawable interface {
	GetPosition() *components.Position
	GetSprite() *components.Sprite
}

type Drawer struct{}

var _ engine.System = new(Drawer)

func (d *Drawer) Update(engine.Entity) error {
	return nil
}

func (d *Drawer) Draw(obj engine.Entity, screen *ebiten.Image) {
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
