package game2

import "github.com/hajimehoshi/ebiten/v2"

type EventName string

const (
	EventNameGameStart            EventName = "gameStart"
	EventNameGameEnd              EventName = "gameEnd"
	EventNameCreate               EventName = "create"
	EventNameDestroy              EventName = "destroy"
	EventNameMouseLeftPress       EventName = "mouseLeftPress"
	EventNameMouseLeftJustPress   EventName = "mouseLeftJustPress"
	EventNameMouseLeftJustRelease EventName = "mouseLeftJustRelease"
	EventNameDraw                 EventName = "draw"
	EventNameUpdate               EventName = "update"
)

type Event interface {
	Name() EventName
}

type EventManager interface {
	On(event EventName, action EventListener)
	Emit(eventName EventName, f func() Event)
}

type eventManager struct {
	eventMap map[EventName][]EventListener
}

func (em *eventManager) On(event EventName, action EventListener) {
	if em.eventMap == nil {
		em.eventMap = make(map[EventName][]EventListener)
	}

	if _, ok := em.eventMap[event]; !ok {
		em.eventMap[event] = make([]EventListener, 0)
	}

	em.eventMap[event] = append(em.eventMap[event], action)
}

func (em *eventManager) Emit(eventName EventName, f func() Event) {
	if listeners, ok := em.eventMap[eventName]; ok {
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

type EventListener func(event Event)

type EventMouseLeftPress struct {
	MouseX, MouseY int
}

type EventGameStart struct{}

func (event EventGameStart) Name() EventName {
	return EventNameGameStart
}

type EventGameEnd struct{}

func (event EventGameEnd) Name() EventName {
	return EventNameGameEnd
}

type EventCreate struct {
	Object Object
}

func (event EventCreate) Name() EventName {
	return EventNameCreate
}

type EventDestroy struct {
	Object Object
}

func (event EventDestroy) Name() EventName {
	return EventNameDestroy
}

func (event EventMouseLeftPress) Name() EventName {
	return EventNameMouseLeftPress
}

type EventMouseLeftJustPress struct {
	MouseX, MouseY int
}

func (event EventMouseLeftJustPress) Name() EventName {
	return EventNameMouseLeftJustPress
}

type EventMouseLeftJustRelease struct {
	MouseX, MouseY int
}

func (event EventMouseLeftJustRelease) Name() EventName {
	return EventNameMouseLeftJustRelease
}

type EventDraw struct {
	Screen *ebiten.Image
}

func (event EventDraw) Name() EventName {
	return EventNameDraw
}

type EventUpdate struct{}

func (event EventUpdate) Name() EventName {
	return EventNameUpdate
}
