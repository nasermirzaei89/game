package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/game/internal/game2"
	"github.com/pkg/errors"
)

func main() {
	if err := ebiten.RunGame(game2.NewGame()); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
