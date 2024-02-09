package events

import "time"

type TestEvent struct {
	name    string
	payload interface{}
}

func (e *TestEvent) Name() string {
	return e.name
}

func (e *TestEvent) OcurredAt() time.Time {
	return time.Now()
}

func (e *TestEvent) Payload() interface{} {
	return e.payload
}
