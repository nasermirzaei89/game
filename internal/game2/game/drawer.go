package game

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/game/internal/game2/engine"
	"image"
	"image/color"
	"log"
)

type drawer struct {
	engine.BaseObject
	bound  image.Rectangle
	active bool
}

func NewDrawer() *drawer {
	d := new(drawer)

	d.OnMouseGlobalLeftPressed(func(ev *engine.EventMouseGlobalLeftPressed) {
		d.bound.Min.X = ev.MouseX
		d.bound.Min.Y = ev.MouseY

		d.active = true
	})

	d.OnMouseGlobalLeftDown(func(ev *engine.EventMouseGlobalLeftDown) {
		d.bound.Max.X = ev.MouseX
		d.bound.Max.Y = ev.MouseY
	})

	d.OnMouseGlobalLeftReleased(func(ev *engine.EventMouseGlobalLeftReleased) {
		d.active = false
	})

	d.OnDraw(func(ev *engine.EventDraw) {
		if !d.active {
			return
		}

		green := color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 128,
		}

		ebitenutil.DrawLine(ev.Screen, float64(d.bound.Min.X), float64(d.bound.Min.Y), float64(d.bound.Max.X), float64(d.bound.Min.Y), green)
		ebitenutil.DrawLine(ev.Screen, float64(d.bound.Min.X), float64(d.bound.Max.Y), float64(d.bound.Max.X), float64(d.bound.Max.Y), green)
		ebitenutil.DrawLine(ev.Screen, float64(d.bound.Min.X), float64(d.bound.Min.Y), float64(d.bound.Min.X), float64(d.bound.Max.Y), green)
		ebitenutil.DrawLine(ev.Screen, float64(d.bound.Max.X), float64(d.bound.Min.Y), float64(d.bound.Max.X), float64(d.bound.Max.Y), green)
	})

	d.OnCreate(func(ev *engine.EventCreate) {
		log.Println(ev.Name())
	})

	d.OnDestroy(func(ev *engine.EventDestroy) {
		log.Println(ev.Name())
	})

	return d
}
