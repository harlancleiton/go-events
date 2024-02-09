package events

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestEventDispatcherSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *ConcreteEventDispatcher
}

func (suite *TestEventDispatcherSuite) SetupTest() {
	suite.dispatcher = NewConcreteEventDispatcher()
	suite.event = TestEvent{name: "Test Event", payload: "Test Payload"}
	suite.event2 = TestEvent{name: "Test Event 2", payload: "Test Payload 2"}
	suite.handler = TestEventHandler{ID: "Handler 1"}
	suite.handler2 = TestEventHandler{ID: "Handler 2"}
	suite.handler3 = TestEventHandler{ID: "Handler 3"}
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Register() {
	err := suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.Name()]))
	suite.Equal(&suite.handler, suite.dispatcher.handlers[suite.event.Name()][0])

	err = suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.Name()]))
	suite.Equal(&suite.handler2, suite.dispatcher.handlers[suite.event.Name()][1])
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Register_WithSameHandler() {
	suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler)
	err := suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler)
	suite.NotNil(err)
	suite.Equal(err, ErrHandlerAlreadyRegistered)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Unregister() {
	suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler)
	suite.dispatcher.RegisterHandler(suite.event2.Name(), &suite.handler2)
	suite.dispatcher.RegisterHandler(suite.event2.Name(), &suite.handler3)

	err := suite.dispatcher.UnregisterHandler(suite.event.Name(), &suite.handler)
	suite.Nil(err)
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.Name()]))

	err = suite.dispatcher.UnregisterHandler(suite.event2.Name(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.Name()]))
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Unregister_WithNonExistentHandler() {
	suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler)
	err := suite.dispatcher.UnregisterHandler(suite.event.Name(), &suite.handler2)
	suite.NotNil(err)
	suite.Equal(err, ErrHandlerNotRegistered)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Clear() {
	suite.Equal(0, len(suite.dispatcher.handlers))
	suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler)
	suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler2)
	suite.dispatcher.RegisterHandler(suite.event2.Name(), &suite.handler3)

	suite.Equal(2, len(suite.dispatcher.handlers))
	suite.dispatcher.Clear()
	suite.Equal(0, len(suite.dispatcher.handlers))
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Has() {
	suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler)
	suite.dispatcher.RegisterHandler(suite.event.Name(), &suite.handler2)

	suite.True(suite.dispatcher.Has(suite.event.Name(), &suite.handler))
	suite.True(suite.dispatcher.Has(suite.event.Name(), &suite.handler2))
	suite.False(suite.dispatcher.Has(suite.event2.Name(), &suite.handler))
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	suite.dispatcher.RegisterHandler(suite.event.Name(), eh)
	suite.dispatcher.Dispatch(&suite.event)

	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh.MethodCalled("Handle", &suite.event)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Dispatch_WithMultipleHandlers() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event)

	suite.dispatcher.RegisterHandler(suite.event.Name(), eh)
	suite.dispatcher.RegisterHandler(suite.event.Name(), eh2)
	suite.dispatcher.Dispatch(&suite.event)

	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)

	eh2.AssertExpectations(suite.T())
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Dispatch_WithNoHandlers() {
	err := suite.dispatcher.Dispatch(&suite.event)
	suite.NotNil(err)
	suite.Equal(err, ErrHandlerNotFound)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestEventDispatcherSuite))
}
