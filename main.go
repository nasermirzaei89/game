package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
)

type game struct{}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(*ebiten.Image) {
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game1 := new(game)

	if err := ebiten.RunGame(game1); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
