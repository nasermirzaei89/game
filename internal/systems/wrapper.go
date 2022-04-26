package systems

import (
	"github.com/nasermirzaei89/game/internal/components"
	"github.com/nasermirzaei89/game/internal/engine"
	"math"
)

type Wrapable interface {
	GetPosition() *components.Position
	GetWrapable() *components.Wrapable
}

type Wrapper struct {
	project *engine.Project
}

var _ engine.System = new(Wrapper)

func (w *Wrapper) Register(p *engine.Project) {
	w.project = p
}

func (w *Wrapper) Update(obj engine.Entity) error {
	entity, ok := obj.(Wrapable)
	if !ok {
		return nil
	}

	position := entity.GetPosition()
	wrapable := entity.GetWrapable()

	width, height := w.project.ScreenWidth, w.project.ScreenHeight

	if wrapable.Horizontal {
		position.X = math.Mod(position.X+float64(width), float64(width))
	}

	if wrapable.Vertical {
		position.Y = math.Mod(position.Y+float64(height), float64(height))
	}

	return nil
}
