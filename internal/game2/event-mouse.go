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

type MouseEventManager interface {
	OnMouseGlobalLeftDown(action func(*EventMouseGlobalLeftDown))
	OnMouseGlobalRightDown(action func(*EventMouseGlobalRightDown))
	OnMouseGlobalMiddleDown(action func(*EventMouseGlobalMiddleDown))
	OnMouseGlobalLeftPressed(action func(*EventMouseGlobalLeftPressed))
	OnMouseGlobalRightPressed(action func(*EventMouseGlobalRightPressed))
	OnMouseGlobalMiddlePressed(action func(*EventMouseGlobalMiddlePressed))
	OnMouseGlobalLeftReleased(action func(*EventMouseGlobalLeftReleased))
	OnMouseGlobalRightReleased(action func(*EventMouseGlobalRightReleased))
	OnMouseGlobalMiddleReleased(action func(*EventMouseGlobalMiddleReleased))
}

type EventMouseGlobalLeftDown struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalLeftDown) Name() EventName {
	return EventNameMouseGlobalLeftDown
}

func (em *eventManager) OnMouseGlobalLeftDown(action func(*EventMouseGlobalLeftDown)) {
	em.On(EventNameMouseGlobalLeftDown, func(event Event) {
		ev, _ := event.(*EventMouseGlobalLeftDown)

		action(ev)
	})
}

type EventMouseGlobalRightDown struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalRightDown) Name() EventName {
	return EventNameMouseGlobalRightDown
}

func (em *eventManager) OnMouseGlobalRightDown(action func(*EventMouseGlobalRightDown)) {
	em.On(EventNameMouseGlobalRightDown, func(event Event) {
		ev, _ := event.(*EventMouseGlobalRightDown)

		action(ev)
	})
}

type EventMouseGlobalMiddleDown struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalMiddleDown) Name() EventName {
	return EventNameMouseGlobalMiddleDown
}

func (em *eventManager) OnMouseGlobalMiddleDown(action func(*EventMouseGlobalMiddleDown)) {
	em.On(EventNameMouseGlobalMiddleDown, func(event Event) {
		ev, _ := event.(*EventMouseGlobalMiddleDown)

		action(ev)
	})
}

type EventMouseGlobalLeftPressed struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalLeftPressed) Name() EventName {
	return EventNameMouseGlobalLeftPressed
}

func (em *eventManager) OnMouseGlobalLeftPressed(action func(*EventMouseGlobalLeftPressed)) {
	em.On(EventNameMouseGlobalLeftPressed, func(event Event) {
		ev, _ := event.(*EventMouseGlobalLeftPressed)

		action(ev)
	})
}

type EventMouseGlobalRightPressed struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalRightPressed) Name() EventName {
	return EventNameMouseGlobalRightPressed
}

func (em *eventManager) OnMouseGlobalRightPressed(action func(*EventMouseGlobalRightPressed)) {
	em.On(EventNameMouseGlobalRightPressed, func(event Event) {
		ev, _ := event.(*EventMouseGlobalRightPressed)

		action(ev)
	})
}

type EventMouseGlobalMiddlePressed struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalMiddlePressed) Name() EventName {
	return EventNameMouseGlobalMiddlePressed
}

func (em *eventManager) OnMouseGlobalMiddlePressed(action func(*EventMouseGlobalMiddlePressed)) {
	em.On(EventNameMouseGlobalMiddlePressed, func(event Event) {
		ev, _ := event.(*EventMouseGlobalMiddlePressed)

		action(ev)
	})
}

type EventMouseGlobalLeftReleased struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalLeftReleased) Name() EventName {
	return EventNameMouseGlobalLeftReleased
}

func (em *eventManager) OnMouseGlobalLeftReleased(action func(*EventMouseGlobalLeftReleased)) {
	em.On(EventNameMouseGlobalLeftReleased, func(event Event) {
		ev, _ := event.(*EventMouseGlobalLeftReleased)

		action(ev)
	})
}

type EventMouseGlobalRightReleased struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalRightReleased) Name() EventName {
	return EventNameMouseGlobalRightReleased
}

func (em *eventManager) OnMouseGlobalRightReleased(action func(*EventMouseGlobalRightReleased)) {
	em.On(EventNameMouseGlobalRightReleased, func(event Event) {
		ev, _ := event.(*EventMouseGlobalRightReleased)

		action(ev)
	})
}

type EventMouseGlobalMiddleReleased struct {
	MouseX, MouseY int
}

func (event EventMouseGlobalMiddleReleased) Name() EventName {
	return EventNameMouseGlobalMiddleReleased
}

func (em *eventManager) OnMouseGlobalMiddleReleased(action func(*EventMouseGlobalMiddleReleased)) {
	em.On(EventNameMouseGlobalMiddleReleased, func(event Event) {
		ev, _ := event.(*EventMouseGlobalMiddleReleased)

		action(ev)
	})
}
