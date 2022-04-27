package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/game/internal/game2"
	"math/rand"
	"time"
)

type batBrain struct {
	game2.BaseObject

	x, y float64
}

func NewBatBrain(g game2.Game) *batBrain {
	bb := new(batBrain)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	bb.OnCreate(func(ev *game2.EventCreate) {
		bb.x = float64(rnd.Intn(g.ScreenWidth()))
		bb.y = float64(rnd.Intn(g.ScreenHeight()))
	})

	bb.OnDraw(func(ev *game2.EventDraw) {
		ebitenutil.DebugPrintAt(ev.Screen, "X", int(bb.x), int(bb.y))
	})

	bb.OnKeyDownLeft(func(ev *game2.EventKeyDownLeft) {
		bb.x -= 10
	})

	bb.OnKeyDownRight(func(ev *game2.EventKeyDownRight) {
		bb.x += 10
	})

	bb.OnKeyDownUp(func(ev *game2.EventKeyDownUp) {
		bb.y -= 10
	})

	bb.OnKeyDownDown(func(ev *game2.EventKeyDownDown) {
		bb.y += 10
	})

	return bb
}
