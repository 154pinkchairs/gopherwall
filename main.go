package main

import (
<<<<<<< HEAD
=======
	"fmt"

>>>>>>> 7b1bb17 (update go.mod)
	"github.com/asynkron/protoactor-go/actor"
	log "github.com/sirupsen/logrus"
)

<<<<<<< HEAD
//building an actor model based firewall's main file
//
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
=======
func main() {
	system := actor.NewActorSystem()
	defer system.Shutdown()

	//Set up actor to receive msgs from submodules
	props := actor.PropsFromFunc(func(context actor.Context) {
		switch msg := context.Message().(type) {
		case string:
			fmt.Println(msg)
		default:
			log.Printf("received unknown message: %v", msg)
		}
	})
	//create actor from props
	receiver := system.Root.Spawn(props)

	//TODO: set up submods to send messages to actor

	//wait for messages
	system.AwaitTermination()
>>>>>>> 7b1bb17 (update go.mod)
}
