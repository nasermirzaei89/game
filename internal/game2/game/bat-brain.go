package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/game/internal/game2/engine"
	"image"
	"math/rand"
	"time"
)

type batBrain struct {
	engine.BaseObject

	x, y float64

	speed float64

	animation engine.Animation
}

func NewBatBrain(g engine.Game) *batBrain {
	bb := new(batBrain)

	bb.speed = 4
	bb.animation = engine.Animation{
		FrameIndex: 0,
		FrameRate:  8,
		Frames: []*engine.Sprite{
			{
				Image: g.MustGetImage("bat-brain-0"),
				Origin: image.Point{
					X: 20,
					Y: 20,
				},
			},
			{
				Image: g.MustGetImage("bat-brain-1"),
				Origin: image.Point{
					X: 20,
					Y: 20,
				},
			},
			{
				Image: g.MustGetImage("bat-brain-2"),
				Origin: image.Point{
					X: 20,
					Y: 20,
				},
			},
		},
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	bb.OnCreate(func(ev *engine.EventCreate) {
		bb.x = float64(rnd.Intn(g.ScreenWidth()))
		bb.y = float64(rnd.Intn(g.ScreenHeight()))
	})

	bb.OnUpdate(func(ev *engine.EventUpdate) {
		bb.animation.Update()
	})

	bb.OnDraw(func(ev *engine.EventDraw) {
		opts := &ebiten.DrawImageOptions{}

		spr := bb.animation.CurrentFrame()

		opts.GeoM.Translate(float64(-spr.Origin.X), float64(-spr.Origin.Y))

		opts.GeoM.Translate(bb.x, bb.y)

		ev.Screen.DrawImage(spr.Image, opts)
	})

	bb.OnKeyDownLeft(func(ev *engine.EventKeyDownLeft) {
		bb.x -= bb.speed
		if bb.x < -20 {
			bb.x = float64(g.ScreenWidth()) + 20
		}
	})

	bb.OnKeyDownRight(func(ev *engine.EventKeyDownRight) {
		bb.x += bb.speed
		if bb.x > float64(g.ScreenWidth())+20 {
			bb.x = -20
		}
	})

	bb.OnKeyDownUp(func(ev *engine.EventKeyDownUp) {
		bb.y -= bb.speed
		if bb.y < -20 {
			bb.y = float64(g.ScreenHeight()) + 20
		}
	})

	bb.OnKeyDownDown(func(ev *engine.EventKeyDownDown) {
		bb.y += bb.speed
		if bb.y > float64(g.ScreenHeight())+20 {
			bb.y = -20
		}
	})

	return bb
}
