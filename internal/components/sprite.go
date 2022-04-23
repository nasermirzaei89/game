package components

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Source           *ebiten.Image
	X0, X1, Y0, Y1   int
	OffsetX, OffsetY int
}

func (s *Sprite) GetSprite() *Sprite {
	return s
}

func (s *Sprite) Image() *ebiten.Image {
	return ebiten.NewImageFromImage(s.Source.SubImage(image.Rect(s.X0, s.Y0, s.X1, s.Y1)))
}
