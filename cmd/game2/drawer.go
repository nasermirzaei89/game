package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nasermirzaei89/game/internal/game2"
	"image"
	"image/color"
	"log"
)

type drawer struct {
	game2.BaseObject
	bound  image.Rectangle
	active bool
}

func NewDrawer() *drawer {
	d := new(drawer)

	d.OnMouseGlobalLeftPressed(func(ev *game2.EventMouseGlobalLeftPressed) {
		d.bound.Min.X = ev.MouseX
		d.bound.Min.Y = ev.MouseY

		d.active = true
	})

	d.OnMouseGlobalLeftDown(func(ev *game2.EventMouseGlobalLeftDown) {
		d.bound.Max.X = ev.MouseX
		d.bound.Max.Y = ev.MouseY
	})

	d.OnMouseGlobalLeftReleased(func(ev *game2.EventMouseGlobalLeftReleased) {
		d.active = false
	})

	d.OnDraw(func(ev *game2.EventDraw) {
		if !d.active {
			return
		}

		ebitenutil.DrawRect(ev.Screen, float64(d.bound.Min.X), float64(d.bound.Min.Y), float64(d.bound.Size().X), float64(d.bound.Size().Y), color.White)
	})

	d.OnCreate(func(ev *game2.EventCreate) {
		log.Println(ev.Name())
	})

	d.OnDestroy(func(ev *game2.EventDestroy) {
		log.Println(ev.Name())
	})

	return d
}
