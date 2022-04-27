package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/game/internal/game2/engine"
	"image"
	"math/rand"
	"time"
)

type orbinaut struct {
	engine.BaseObject

	x, y           float64
	speedX, speedY float64

	animation engine.Animation
}

func NewOrbinaut(g engine.Game) *orbinaut {
	o := new(orbinaut)

	o.animation = engine.Animation{
		FrameIndex: 0,
		FrameRate:  16,
		Frames: []*engine.Sprite{
			{
				Image: g.MustGetImage("orbinaut-0"),
				Origin: image.Point{
					X: 24,
					Y: 24,
				},
			},
			{
				Image: g.MustGetImage("orbinaut-1"),
				Origin: image.Point{
					X: 24,
					Y: 24,
				},
			},
			{
				Image: g.MustGetImage("orbinaut-2"),
				Origin: image.Point{
					X: 24,
					Y: 24,
				},
			},
		},
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	o.OnCreate(func(ev *engine.EventCreate) {
		o.x = float64(rnd.Intn(g.ScreenWidth()))
		o.y = float64(rnd.Intn(g.ScreenHeight()))

		o.speedX = float64(rnd.Intn(8)) - 4
		o.speedY = float64(rnd.Intn(8)) - 4
	})

	o.OnUpdate(func(ev *engine.EventUpdate) {
		o.animation.Update()

		o.x += o.speedX
		o.y += o.speedY

		if o.x < -20 {
			o.x = float64(g.ScreenWidth()) + 20
		}
		if o.x > float64(g.ScreenWidth())+20 {
			o.x = -20
		}
		if o.y < -20 {
			o.y = float64(g.ScreenHeight()) + 20
		}
		if o.y > float64(g.ScreenHeight())+20 {
			o.y = -20
		}
	})

	o.OnDraw(func(ev *engine.EventDraw) {
		opts := &ebiten.DrawImageOptions{}

		spr := o.animation.CurrentFrame()

		opts.GeoM.Translate(float64(-spr.Origin.X), float64(-spr.Origin.Y))

		opts.GeoM.Translate(o.x, o.y)

		ev.Screen.DrawImage(spr.Image, opts)
	})

	return o
}
