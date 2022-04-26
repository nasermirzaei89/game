package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image            *ebiten.Image
	OffsetX, OffsetY int
}

func (s *Sprite) GetSprite() *Sprite {
	return s
}
