package game2

const (
	EventNameCreate  EventName = "create"
	EventNameDestroy EventName = "destroy"
)

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
