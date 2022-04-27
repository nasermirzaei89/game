package game2

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pkg/errors"
)

type Game interface {
	Run() error
	On(event EventName, action EventListener)
	OnDraw(action func(ev *EventDraw))
	OnUpdate(action func())
}

type game struct {
	eventMap map[EventName][]EventListener
	started  bool
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
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *game) On(event EventName, action EventListener) {
	if _, ok := g.eventMap[event]; !ok {
		g.eventMap[event] = make([]EventListener, 0)
	}

	g.eventMap[event] = append(g.eventMap[event], action)
}

func (g *game) OnDraw(action func(ev *EventDraw)) {
	g.On(EventNameDraw, func(event Event) {
		eventDraw := event.(*EventDraw)

		action(eventDraw)
	})
}

func (g *game) OnUpdate(action func()) {
	g.On(EventNameUpdate, func(event Event) {
		action()
	})
}

func (g *game) Emit(eventName EventName, f func() Event) {
	if listeners, ok := g.eventMap[eventName]; ok {
		if len(listeners) == 0 {
			return
		}

		ev := f()
		if ev == nil {
			return
		}

		for i := range listeners {
			listeners[i](ev)
		}
	}
}

func NewGame() Game {
	g := &game{
		eventMap: make(map[EventName][]EventListener),
		started:  false,
	}

	return g
}
