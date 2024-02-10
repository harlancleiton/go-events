package events

import (
	"fmt"
	"sync"

	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event Event, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

type TestEventHandler struct {
	ID string
}

func (h *TestEventHandler) Handle(event Event, wg *sync.WaitGroup) {
	fmt.Println("TestEventHandler: ", h.ID, event.Name())
	wg.Done()
}
