package game2

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pkg/errors"
)

type Game interface {
	Run() error

	EventManager

	AddObject(obj Object)
	RemoveObject(obj Object)

	HierarchicalEmit(eventName EventName, f func() Event)

	OnStart(action func(ev *EventGameStart))
	OnEnd(action func(ev *EventGameEnd))

	ScreenWidth() int
	ScreenHeight() int
}

type game struct {
	eventManager
	buffer  []Object
	objects []Object
	started bool

	screenWidth, screenHeight int
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
	g.HierarchicalEmit(EventNameUpdate, func() Event {
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

	g.HierarchicalEmit(EventNameMouseGlobalLeftDown, func() Event {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalLeftDown{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalRightDown, func() Event {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalRightDown{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalMiddleDown, func() Event {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalMiddleDown{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalLeftPressed, func() Event {
		if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalLeftPressed{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalRightPressed, func() Event {
		if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalRightPressed{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalMiddlePressed, func() Event {
		if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalMiddlePressed{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalLeftReleased, func() Event {
		if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalLeftReleased{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalRightReleased, func() Event {
		if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalRightReleased{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameMouseGlobalMiddleReleased, func() Event {
		if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonMiddle) {
			return nil
		}

		x, y := ebiten.CursorPosition()

		return &EventMouseGlobalMiddleReleased{
			MouseX: x,
			MouseY: y,
		}
	})

	g.HierarchicalEmit(EventNameKeyDownLeft, func() Event {
		if !ebiten.IsKeyPressed(ebiten.KeyLeft) {
			return nil
		}

		return &EventKeyDownLeft{}
	})

	g.HierarchicalEmit(EventNameKeyDownRight, func() Event {
		if !ebiten.IsKeyPressed(ebiten.KeyRight) {
			return nil
		}

		return &EventKeyDownRight{}
	})

	g.HierarchicalEmit(EventNameKeyDownUp, func() Event {
		if !ebiten.IsKeyPressed(ebiten.KeyUp) {
			return nil
		}

		return &EventKeyDownUp{}
	})

	g.HierarchicalEmit(EventNameKeyDownDown, func() Event {
		if !ebiten.IsKeyPressed(ebiten.KeyDown) {
			return nil
		}

		return &EventKeyDownDown{}
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
	g.HierarchicalEmit(EventNameDraw, func() Event {
		return &EventDraw{Screen: screen}
	})
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.screenWidth = outsideWidth
	g.screenHeight = outsideHeight

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

func (g *game) HierarchicalEmit(eventName EventName, f func() Event) {
	g.Emit(eventName, f)

	for i := range g.objects {
		g.objects[i].Emit(eventName, f)
	}
}

func (g *game) ScreenWidth() int {
	return g.screenWidth
}

func (g *game) ScreenHeight() int {
	return g.screenHeight
}

func (g *game) OnStart(action func(*EventGameStart)) {
	g.On(EventNameGameStart, func(event Event) {
		eventGameStart, _ := event.(*EventGameStart)

		action(eventGameStart)
	})
}

func (g *game) OnEnd(action func(*EventGameEnd)) {
	g.On(EventNameGameEnd, func(event Event) {
		eventGameEnd, _ := event.(*EventGameEnd)

		action(eventGameEnd)
	})
}

func NewGame() Game {
	g := &game{
		eventManager: eventManager{},
		buffer:       make([]Object, 0),
		objects:      make([]Object, 0),
		started:      false,
		screenWidth:  0,
		screenHeight: 0,
	}

	return g
}
