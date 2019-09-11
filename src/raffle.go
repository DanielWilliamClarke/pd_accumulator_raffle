package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Equanox/gotron"
)

// Create a custom event struct that has a pointer to gotron.Event
type CustomEvent struct {
	*gotron.Event
	CustomAttribute string `json:"AtrNameInFrontend"`
}

func main() {

	// Create a new browser window instance
	window, err := gotron.New("ui/build")
	if err != nil {
		log.Println(err)
		return
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 600
	window.WindowOptions.Title = "Gotron"

	done, err := window.Start()
	if err != nil {
		log.Println(err)
		return
	}

	onEvent := gotron.Event{Event: "hello-back"}
	window.On(&onEvent, func(bin []byte) {

		log.Println("received hello")
		log.Println(string(bin))

		window.Send(&CustomEvent{
			Event:           &gotron.Event{Event: "hello-front"},
			CustomAttribute: fmt.Sprintf("Hello frontend! - %s", time.Now()),
		})
	})

	//window.OpenDevTools()

	<-done

	// read input file of names

	// for each time interval pick a name from the list using some kind of random number generator
	// add a point to that list

	// show timer until next pick
	// show picked name
	// show live scorebord

}
