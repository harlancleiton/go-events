package events

import "time"

type Event interface {
	Name() string
	OcurredAt() time.Time
	Payload() interface{}
}
