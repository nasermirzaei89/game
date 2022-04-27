package game2

import "github.com/hajimehoshi/ebiten/v2"

type EventName string

const (
	EventNameGameStart            EventName = "gameStart"
	EventNameGameEnd              EventName = "gameEnd"
	EventNameMouseLeftPress       EventName = "mouseLeftPress"
	EventNameMouseLeftJustPress   EventName = "mouseLeftJustPress"
	EventNameMouseLeftJustRelease EventName = "mouseLeftJustRelease"
	EventNameDraw                 EventName = "draw"
	EventNameUpdate               EventName = "update"
)

type Event interface {
	Name() EventName
}

type EventListener func(event Event)

type EventMouseLeftPress struct {
	MouseX, MouseY int
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

type EventGameStart struct{}

func (event EventGameStart) Name() EventName {
	return EventNameGameStart
}

type EventGameEnd struct{}

func (event EventGameEnd) Name() EventName {
	return EventNameGameEnd
}
