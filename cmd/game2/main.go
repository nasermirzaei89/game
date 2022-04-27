package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/game/internal/game2"
	"github.com/pkg/errors"
)

func main() {
	g := game2.NewGame()

	d := NewDrawer()

	g.AddObject(d)

	frame := 0

	g.OnStart(func(ev *game2.EventGameStart) {
		log.Println(ev.Name())
	})

	g.OnEnd(func(ev *game2.EventGameEnd) {
		log.Println(ev.Name())
	})

	g.OnUpdate(func(ev *game2.EventUpdate) {
		frame++

		if frame == 100 {
			g.RemoveObject(d)
		}
	})

	g.OnDraw(func(ev *game2.EventDraw) {
		ebitenutil.DebugPrint(ev.Screen, fmt.Sprintf("FPS: %f\nFrame: %d", ebiten.CurrentFPS(), frame))
	})

	g.AddObject(NewBatBrain(g))

	if err := g.Run(); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
