package engine

import "github.com/hajimehoshi/ebiten/v2"

const (
	EventNameGameStart EventName = "gameStart"
	EventNameGameEnd   EventName = "gameEnd"

	EventNameDraw   EventName = "draw"
	EventNameUpdate EventName = "update"
)

type EventGameStart struct{}

func (event EventGameStart) Name() EventName {
	return EventNameGameStart
}

type EventGameEnd struct{}

func (event EventGameEnd) Name() EventName {
	return EventNameGameEnd
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
