package game2

const (
	EventNameKeyDownLeft  EventName = "keyDownLeft"
	EventNameKeyDownRight EventName = "keyDownRight"
	EventNameKeyDownUp    EventName = "keyDownUp"
	EventNameKeyDownDown  EventName = "keyDownDown"

	EventNameKeyPressedLeft  EventName = "keyPressedLeft"
	EventNameKeyPressedRight EventName = "keyPressedRight"
	EventNameKeyPressedUp    EventName = "keyPressedUp"
	EventNameKeyPressedDown  EventName = "keyPressedDown"

	EventNameKeyReleasedLeft  EventName = "keyReleasedLeft"
	EventNameKeyReleasedRight EventName = "keyReleasedRight"
	EventNameKeyReleasedUp    EventName = "keyReleasedUp"
	EventNameKeyReleasedDown  EventName = "keyReleasedDown"
)

type KeyboardEventManager interface {
	OnKeyDownLeft(action func(*EventKeyDownLeft))
	OnKeyDownRight(action func(*EventKeyDownRight))
	OnKeyDownUp(action func(*EventKeyDownUp))
	OnKeyDownDown(action func(*EventKeyDownDown))

	OnKeyPressedLeft(action func(*EventKeyPressedLeft))
	OnKeyPressedRight(action func(*EventKeyPressedRight))
	OnKeyPressedUp(action func(*EventKeyPressedUp))
	OnKeyPressedDown(action func(*EventKeyPressedDown))

	OnKeyReleasedLeft(action func(*EventKeyReleasedLeft))
	OnKeyReleasedRight(action func(*EventKeyReleasedRight))
	OnKeyReleasedUp(action func(*EventKeyReleasedUp))
	OnKeyReleasedDown(action func(*EventKeyReleasedDown))
}

type EventKeyDownLeft struct{}

func (event EventKeyDownLeft) Name() EventName {
	return EventNameKeyDownLeft
}

func (em *eventManager) OnKeyDownLeft(action func(*EventKeyDownLeft)) {
	em.On(EventNameKeyDownLeft, func(event Event) {
		ev, _ := event.(*EventKeyDownLeft)

		action(ev)
	})
}

type EventKeyDownRight struct{}

func (event EventKeyDownRight) Name() EventName {
	return EventNameKeyDownRight
}

func (em *eventManager) OnKeyDownRight(action func(*EventKeyDownRight)) {
	em.On(EventNameKeyDownRight, func(event Event) {
		ev, _ := event.(*EventKeyDownRight)

		action(ev)
	})
}

type EventKeyDownUp struct{}

func (event EventKeyDownUp) Name() EventName {
	return EventNameKeyDownUp
}

func (em *eventManager) OnKeyDownUp(action func(*EventKeyDownUp)) {
	em.On(EventNameKeyDownUp, func(event Event) {
		ev, _ := event.(*EventKeyDownUp)

		action(ev)
	})
}

type EventKeyDownDown struct{}

func (event EventKeyDownDown) Name() EventName {
	return EventNameKeyDownDown
}

func (em *eventManager) OnKeyDownDown(action func(*EventKeyDownDown)) {
	em.On(EventNameKeyDownDown, func(event Event) {
		ev, _ := event.(*EventKeyDownDown)

		action(ev)
	})
}

type EventKeyPressedLeft struct{}

func (event EventKeyPressedLeft) Name() EventName {
	return EventNameKeyPressedLeft
}

func (em *eventManager) OnKeyPressedLeft(action func(*EventKeyPressedLeft)) {
	em.On(EventNameKeyPressedLeft, func(event Event) {
		ev, _ := event.(*EventKeyPressedLeft)

		action(ev)
	})
}

type EventKeyPressedRight struct{}

func (event EventKeyPressedRight) Name() EventName {
	return EventNameKeyPressedRight
}

func (em *eventManager) OnKeyPressedRight(action func(*EventKeyPressedRight)) {
	em.On(EventNameKeyPressedRight, func(event Event) {
		ev, _ := event.(*EventKeyPressedRight)

		action(ev)
	})
}

type EventKeyPressedUp struct{}

func (event EventKeyPressedUp) Name() EventName {
	return EventNameKeyPressedUp
}

func (em *eventManager) OnKeyPressedUp(action func(*EventKeyPressedUp)) {
	em.On(EventNameKeyPressedUp, func(event Event) {
		ev, _ := event.(*EventKeyPressedUp)

		action(ev)
	})
}

type EventKeyPressedDown struct{}

func (event EventKeyPressedDown) Name() EventName {
	return EventNameKeyPressedDown
}

func (em *eventManager) OnKeyPressedDown(action func(*EventKeyPressedDown)) {
	em.On(EventNameKeyPressedDown, func(event Event) {
		ev, _ := event.(*EventKeyPressedDown)

		action(ev)
	})
}

type EventKeyReleasedLeft struct{}

func (event EventKeyReleasedLeft) Name() EventName {
	return EventNameKeyReleasedLeft
}

func (em *eventManager) OnKeyReleasedLeft(action func(*EventKeyReleasedLeft)) {
	em.On(EventNameKeyReleasedLeft, func(event Event) {
		ev, _ := event.(*EventKeyReleasedLeft)

		action(ev)
	})
}

type EventKeyReleasedRight struct{}

func (event EventKeyReleasedRight) Name() EventName {
	return EventNameKeyReleasedRight
}

func (em *eventManager) OnKeyReleasedRight(action func(*EventKeyReleasedRight)) {
	em.On(EventNameKeyReleasedRight, func(event Event) {
		ev, _ := event.(*EventKeyReleasedRight)

		action(ev)
	})
}

type EventKeyReleasedUp struct{}

func (event EventKeyReleasedUp) Name() EventName {
	return EventNameKeyReleasedUp
}

func (em *eventManager) OnKeyReleasedUp(action func(*EventKeyReleasedUp)) {
	em.On(EventNameKeyReleasedUp, func(event Event) {
		ev, _ := event.(*EventKeyReleasedUp)

		action(ev)
	})
}

type EventKeyReleasedDown struct{}

func (event EventKeyReleasedDown) Name() EventName {
	return EventNameKeyReleasedDown
}

func (em *eventManager) OnKeyReleasedDown(action func(*EventKeyReleasedDown)) {
	em.On(EventNameKeyReleasedDown, func(event Event) {
		ev, _ := event.(*EventKeyReleasedDown)

		action(ev)
	})
}
