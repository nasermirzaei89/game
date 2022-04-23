package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
	"math"
)

type Wrapable interface {
	GetPosition() *components.Position
	GetWrapable() *components.Wrapable
}

type Wrapper struct{}

var _ engine.System = new(Wrapper)

func (w *Wrapper) Update(obj engine.Entity) error {
	entity, ok := obj.(Wrapable)
	if !ok {
		return nil
	}

	position := entity.GetPosition()
	wrapable := entity.GetWrapable()

	width, height := ebiten.WindowSize()

	if wrapable.Horizontal {
		position.X = math.Mod(position.X+float64(width), float64(width))
	}

	if wrapable.Vertical {
		position.Y = math.Mod(position.Y+float64(height), float64(height))
	}

	return nil
}

func (w *Wrapper) Draw(obj engine.Entity, screen *ebiten.Image) {}
