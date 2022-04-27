package engine

type Object interface {
	EventManager
}

type BaseObject struct {
	eventManager
}

var _ Object = new(BaseObject)

func (obj *BaseObject) OnDraw(action func(ev *EventDraw)) {
	obj.On(EventNameDraw, func(event Event) {
		eventDraw, _ := event.(*EventDraw)

		action(eventDraw)
	})
}

func (obj *BaseObject) OnCreate(action func(ev *EventCreate)) {
	obj.On(EventNameCreate, func(event Event) {
		eventCreate, _ := event.(*EventCreate)

		action(eventCreate)
	})
}

func (obj *BaseObject) OnDestroy(action func(ev *EventDestroy)) {
	obj.On(EventNameDestroy, func(event Event) {
		eventDestroy, _ := event.(*EventDestroy)

		action(eventDestroy)
	})
}
