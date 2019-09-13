package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/danielwilliamclarke/raffle/accumulator"
)

func main() {

	participantStr, err := ioutil.ReadFile("./participants.txt")
	if err != nil {
		log.Fatalf("No such participants file: %v", err)
	}

	log.Println("Raffle Participants ---------------------------------------------")
	log.Println(string(participantStr))

	participants := strings.Split(string(participantStr), "\n")

	// Shuffle participants
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(participants), func(i, j int) {
		participants[i], participants[j] = participants[j], participants[i]
	})

	// Link participants to scores
	pariticpantScores := accumulator.ParseParticiapants(participants)

	// Create a new browser window instance
	window, done := accumulator.CreateWindow("ui/build")

	// Begin the Raffle!
	raffler := accumulator.NewRaffle(window)
	go raffler.Run(pariticpantScores)

	<-done
}
