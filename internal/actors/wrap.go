package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/td/internal/components"
	"github.com/nasermirzaei89/td/internal/engine"
	"math"
)

type Wrapable interface {
	GetPosition() *components.Position
	IsWrapable()
}

type Wrap struct{}

func (wrap *Wrap) Update(obj engine.Object) error {
	entity, ok := obj.(Wrapable)
	if !ok {
		return nil
	}

	position := entity.GetPosition()

	w, h := ebiten.WindowSize()

	position.X = math.Mod(position.X+float64(w), float64(w))
	position.Y = math.Mod(position.Y+float64(h), float64(h))

	return nil
}

func (wrap *Wrap) Draw(obj engine.Object, screen *ebiten.Image) {}

var _ engine.Actor = new(Wrap)
