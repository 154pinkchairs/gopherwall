package main

import (
	"github.com/asynkron/protoactor-go/actor"
	log "github.com/sirupsen/logrus"
)

// building an actor model based firewall's main file
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
	}
}
