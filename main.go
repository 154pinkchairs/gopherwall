package main

import (
	"github.com/asynkron/protoactor-go/actor"
	log "github.com/sirupsen/logrus"
)

type Firewall struct {
}

func (state *Firewall) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		log.Info("Firewall started")
	case *actor.Stopping:
		log.Info("Firewall stopping")
	case *actor.Stopped:
		log.Info("Firewall stopped")
	case *actor.Restarting:
		log.Info("Firewall restarting")
	default:
		log.Info("Firewall received unknown message")
		_ = msg
}
}
//listen to messages from the actors in the packages
func main() {
	context := actor.NewActorSystem().Root
	props := actor.PropsFromProducer(func() actor.Actor { return &Firewall{} })
	pid, err := context.SpawnNamed(props, "firewall")
	if err != nil {
		log.Fatal(err)
	}
	context.Send(pid, &actor.Started{})
}
