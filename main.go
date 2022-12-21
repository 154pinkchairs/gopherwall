package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	system := actor.NewActorSystem()
	defer system.Terminate()

	//Set up actor to receive msgs from submodules
	props := actor.FromFunc(func(context actor.Context) {
		switch msg := context.Message().(type) {
		case string:
			fmt.Println(msg)
		default:
			log.Printf("received unknown message: %v", msg)
		}
	})
	receiver, err := system.Root.CreateActor(props)
	if err != nil {
		log.Fatalf("error creating actor: %v", err)
	}

	//TODO: set up submods to send messages to actor

	//wait for messages
	system.AwaitTermination()
}
