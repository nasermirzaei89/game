package game2

const (
	EventNameMouseGlobalLeftDown       EventName = "mouseGlobalLeftDown"
	EventNameMouseGlobalRightDown      EventName = "mouseGlobalRightDown"
	EventNameMouseGlobalMiddleDown     EventName = "mouseGlobalMiddleDown"
	EventNameMouseGlobalLeftPressed    EventName = "mouseGlobalLeftPressed"
	EventNameMouseGlobalRightPressed   EventName = "mouseGlobalRightPressed"
	EventNameMouseGlobalMiddlePressed  EventName = "mouseGlobalMiddlePressed"
	EventNameMouseGlobalLeftReleased   EventName = "mouseGlobalLeftReleased"
	EventNameMouseGlobalRightReleased  EventName = "mouseGlobalRightReleased"
	EventNameMouseGlobalMiddleReleased EventName = "mouseGlobalMiddleReleased"
)

type EventMouseGlobalLeftDown struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalLeftDown) Name() EventName {
	return EventNameMouseGlobalLeftDown
}

type EventMouseGlobalRightDown struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalRightDown) Name() EventName {
	return EventNameMouseGlobalRightDown
}

type EventMouseGlobalMiddleDown struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalMiddleDown) Name() EventName {
	return EventNameMouseGlobalMiddleDown
}

type EventMouseGlobalLeftPressed struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalLeftPressed) Name() EventName {
	return EventNameMouseGlobalLeftPressed
}

type EventMouseGlobalRightPressed struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalRightPressed) Name() EventName {
	return EventNameMouseGlobalRightPressed
}

type EventMouseGlobalMiddlePressed struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalMiddlePressed) Name() EventName {
	return EventNameMouseGlobalMiddlePressed
}

type EventMouseGlobalLeftReleased struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalLeftReleased) Name() EventName {
	return EventNameMouseGlobalLeftReleased
}

type EventMouseGlobalRightReleased struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalRightReleased) Name() EventName {
	return EventNameMouseGlobalRightReleased
}

type EventMouseGlobalMiddleReleased struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalMiddleReleased) Name() EventName {
	return EventNameMouseGlobalMiddleReleased
}
