package game2

type EventName string

type Event interface {
	Name() EventName
}

type EventListener func(event Event)

type EventManager interface {
	On(event EventName, action EventListener)
	Emit(eventName EventName, f func() Event)
	ExtendedEventManager
}

type ExtendedEventManager interface {
	OnDraw(action func(*EventDraw))
	OnUpdate(action func(*EventUpdate))
	MouseEventManager
	KeyboardEventManager
}

type eventManager struct {
	eventMap map[EventName][]EventListener
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

func (em *eventManager) On(event EventName, action EventListener) {
	if em.eventMap == nil {
		em.eventMap = make(map[EventName][]EventListener)
	}

	if _, ok := em.eventMap[event]; !ok {
		em.eventMap[event] = make([]EventListener, 0)
	}

	em.eventMap[event] = append(em.eventMap[event], action)
}

func (em *eventManager) OnDraw(action func(*EventDraw)) {
	em.On(EventNameDraw, func(event Event) {
		ev, _ := event.(*EventDraw)

		action(ev)
	})
}

func (em *eventManager) OnUpdate(action func(*EventUpdate)) {
	em.On(EventNameUpdate, func(event Event) {
		ev, _ := event.(*EventUpdate)

		action(ev)
	})
}
