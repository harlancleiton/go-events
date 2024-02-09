package events

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event Event) {
	m.Called(event)
}

type TestEventHandler struct {
	ID string
}

func (h *TestEventHandler) Handle(event Event) {
	fmt.Println("TestEventHandler: ", h.ID, event.Name())
}
