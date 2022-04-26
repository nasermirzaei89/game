package main

import (
	"github.com/nasermirzaei89/game/internal/game1/game"
	"github.com/pkg/errors"
)

func main() {
	if err := game.Run(); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
