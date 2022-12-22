package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockActorSystem struct {
	mock.Mock
}

func (m *mockActorSystem) Terminate() {
	m.Called()
}

func (m *mockActorSystem) Root() *actor.Actor {
	args := m.Called()
	return args.Get(0).(*actor.Actor)
}

func (m *mockActorSystem) CreateActor(props *actor.Props) (*actor.Actor, error) {
	args := m.Called(props)
	return args.Get(0).(*actor.Actor), args.Error(1)
}

func (m *mockActorSystem) AwaitTermination() {
	m.Called()
}

func TestMain(t *testing.T) {
	system := new(mockActorSystem)
	system.On("Terminate")

	rootActor := new(actor.MockActor)
	system.On("Root").Return(rootActor)

	props := actor.FromFunc(func(context actor.Context) {})
	actor := new(actor.MockActor)
	system.On("CreateActor", props).Return(actor, nil)

	system.On("AwaitTermination")

	main(system)

	system.AssertExpectations(t)
	rootActor.AssertExpectations(t)
	actor.AssertExpectations(t)

	//NOTE: that does not cover a scenario, where an err is returned when creating an actor, so we
	//may want to add additional tests for that
}
