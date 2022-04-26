package systems

import (
	components2 "github.com/nasermirzaei89/game/internal/game1/components"
	engine2 "github.com/nasermirzaei89/game/internal/game1/engine"
	"math"
)

type Wrapable interface {
	GetPosition() *components2.Position
	GetWrapable() *components2.Wrapable
}

type Wrapper struct {
	project *engine2.Project
}

var _ engine2.System = new(Wrapper)

func (w *Wrapper) Register(p *engine2.Project) {
	w.project = p
}

func (w *Wrapper) Update(obj engine2.Entity) error {
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
