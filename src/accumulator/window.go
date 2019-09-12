package accumulator

import (
	"log"

	"github.com/Equanox/gotron"
)

func CreateWindow(uiPath string) (*gotron.BrowserWindow, chan bool) {
	window, err := gotron.New(uiPath)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 1000
	window.WindowOptions.Title = "Gotron"

	done, err := window.Start()
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	return window, done
}
