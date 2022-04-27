package game2

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pkg/errors"
)

type Game interface {
	Run() error
	EventManager
	OnDraw(action func(ev *EventDraw))
	OnUpdate(action func())
	AddObject(obj Object)
	RemoveObject(obj Object)
}

type game struct {
	eventManager
	buffer  []Object
	objects []Object
	started bool
}

var _ ebiten.Game = new(game)

var _ Game = new(game)

func (g *game) Run() error {
	if err := ebiten.RunGame(g); err != nil {
		return errors.Wrap(err, "error on run ebiten game")
	}

	return nil
}

func (g *game) Update() error {
	g.Emit(EventNameUpdate, func() Event {
		return &EventUpdate{}
	})

	g.Emit(EventNameGameStart, func() Event {
		if !g.started {
			g.started = true

			return &EventGameStart{}
		}

		return nil
	})

	for i := range g.buffer {
		g.objects = append(g.objects, g.buffer[i])

		g.buffer[i].Emit(EventNameCreate, func() Event {
			return &EventCreate{Object: g.buffer[i]}
		})
	}

	g.buffer = make([]Object, 0)

	g.Emit(EventNameMouseLeftPress, func() Event {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseLeftPress{
			MouseX: x,
			MouseY: y,
		}
	})

	g.Emit(EventNameMouseLeftJustPress, func() Event {
		if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseLeftJustPress{
			MouseX: x,
			MouseY: y,
		}
	})

	g.Emit(EventNameMouseLeftJustRelease, func() Event {
		if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseLeftJustRelease{
			MouseX: x,
			MouseY: y,
		}
	})

	g.Emit(EventNameGameEnd, func() Event {
		if ebiten.IsWindowBeingClosed() {
			return &EventGameEnd{}
		}

		return nil
	})

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.Emit(EventNameDraw, func() Event {
		return &EventDraw{Screen: screen}
	})

	for i := range g.objects {
		g.objects[i].Emit(EventNameDraw, func() Event {
			return &EventDraw{Screen: screen}
		})
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *game) AddObject(obj Object) {
	g.buffer = append(g.buffer, obj)
}

func (g *game) RemoveObject(obj Object) {
	obj.Emit(EventNameDestroy, func() Event {
		return &EventDestroy{Object: obj}
	})

	for i := range g.objects {
		if obj == g.objects[i] {
			g.objects = append(g.objects[:i], g.objects[i+1:]...)

			return
		}
	}
}

func (g *game) OnDraw(action func(ev *EventDraw)) {
	g.On(EventNameDraw, func(event Event) {
		eventDraw, _ := event.(*EventDraw)

		action(eventDraw)
	})
}

func (g *game) OnUpdate(action func()) {
	g.On(EventNameUpdate, func(event Event) {
		action()
	})
}

func NewGame() Game {
	g := &game{
		started: false,
	}

	return g
}
