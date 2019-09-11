package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/danielwilliamclarke/raffle/accumulator"
)

func main() {

	participantStr, err := ioutil.ReadFile("./participants.txt")
	if err != nil {
		log.Fatalf("No such participants file: %v", err)
	}

	participants := strings.Split(string(participantStr), "\n")
	pariticpantScores := accumulator.ParseParticiapants(participants)

	// Create a new browser window instance
	window, done := accumulator.CreateWindow("ui/build")

	raffler := accumulator.Raffler{
		Participants: pariticpantScores,
		Window:       window,
	}
	go raffler.Run()

	<-done
}
