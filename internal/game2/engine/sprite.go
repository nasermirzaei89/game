package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Sprite struct {
	Image  *ebiten.Image
	Origin image.Point
}
