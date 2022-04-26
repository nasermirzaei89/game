package game2

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pkg/errors"
	"log"
)

type Game struct {
	eventMap map[EventName]map[EventListenerID]EventListener
	started  bool
}

type EventName string

const (
	EventNameGameStart   EventName = "gameStart"
	EventNameGameEnd     EventName = "gameEnd"
	EventNameGlobalClick EventName = "globalClick"
)

type EventGlobalClick struct {
	MouseX, MouseY int
}

func (event EventGlobalClick) Name() EventName {
	return EventNameGlobalClick
}

type EventListenerID string

type Event interface {
	Name() EventName
}

type EventListener func(event Event)

func (game *Game) Update() error {
	func() {
		if !game.started {
			game.DispatchEvent(EventNameGameStart, nil)
			game.started = true
		}
	}()

	func() {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			game.DispatchEvent(EventNameGlobalClick, &EventGlobalClick{
				MouseX: x,
				MouseY: y,
			})
		}
	}()

	func() {
		if ebiten.IsWindowBeingClosed() {
			game.DispatchEvent(EventNameGameEnd, nil)
		}
	}()

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (game *Game) AddEventListener(event EventName, action EventListener) EventListenerID {
	eventID := EventListenerID(uuid.NewString())

	if _, ok := game.eventMap[event]; !ok {
		game.eventMap[event] = make(map[EventListenerID]EventListener)
	}

	game.eventMap[event][eventID] = action

	return eventID
}

func (game *Game) RemoveEventListener(event EventName, id EventListenerID) error {
	if _, ok := game.eventMap[event]; !ok {
		return errors.Errorf("event '%s' doesn't exists", event)
	}

	if _, ok := game.eventMap[event][id]; !ok {
		return errors.Errorf("event listener with id '%s' on event '%s' doesn't exists", id, event)
	}

	delete(game.eventMap[event], id)

	return nil
}

func (game *Game) DispatchEvent(eventName EventName, event Event) {
	if listeners, ok := game.eventMap[eventName]; ok {
		for _, f := range listeners {
			f(event)
		}
	}
}

var _ ebiten.Game = new(Game)

func NewGame() *Game {
	game := &Game{
		eventMap: make(map[EventName]map[EventListenerID]EventListener),
		started:  false,
	}

	game.AddEventListener(EventNameGameStart, func(_ Event) {
		log.Println("Game Started!")
	})

	game.AddEventListener(EventNameGameEnd, func(_ Event) {
		log.Println("Game Ended!")
	})

	game.AddEventListener(EventNameGlobalClick, func(ev Event) {
		eventGlobalClick := ev.(*EventGlobalClick)

		log.Println("Clicked!", eventGlobalClick.MouseX, eventGlobalClick.MouseY)
	})

	id := game.AddEventListener(EventNameGlobalClick, func(_ Event) {
		log.Println("Yeah!")
	})

	_ = game.RemoveEventListener(EventNameGlobalClick, id)

	return game
}
