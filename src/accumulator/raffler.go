package accumulator

import (
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/Equanox/gotron"
)

type raffleUpdate struct {
	*gotron.Event
	Update raffleUpdateAttributes `json:"update"`
}

type raffleUpdateAttributes struct {
	Round               int                 `json:"round"`
	GoldenRound         goldenRound         `json:"goldenRound"`
	SelectedParticipant PartiticpantScore   `json:"selectedParticipant"`
	ScoreBoard          []PartiticpantScore `json:"scoreBoard"`
}

type goldenRound struct {
	Next bool `json:"next"`
	Now  bool `json:"now"`
}

// PartiticpantScore links pariticpant names to a score
type PartiticpantScore struct {
	Participant string `json:"participant"`
	Score       int    `json:"score"`
}

type raffleTimer struct {
	elaspedTime     time.Duration
	regularInterval time.Duration
	goldenInterval  time.Duration
}

// ParseParticiapants transforms array names into partiticpantScore structs
func ParseParticiapants(participants []string) (ps []PartiticpantScore) {
	for _, p := range participants {
		ps = append(ps, PartiticpantScore{
			Participant: p,
			Score:       0,
		})
	}
	return ps
}

// NewRaffle raffle constructor
func NewRaffle(window *gotron.BrowserWindow) Raffler {
	return Raffler{
		window: window,
		timings: &raffleTimer{
			elaspedTime:     0 * time.Second,
			regularInterval: 20 * time.Second,
			goldenInterval:  300 * time.Second,
		},
	}
}

// Raffler provides functionality to run raffle accumulator
type Raffler struct {
	window  *gotron.BrowserWindow
	timings *raffleTimer
	round   int
}

// Run runs raffle loop
func (r Raffler) Run(participants []PartiticpantScore) {

	log.Println("Starting Accumulator Raffle")
	rand.Seed(time.Now().UnixNano())

	for {

		time.Sleep(r.timings.regularInterval)

		// Determine score and golden round
		r.round++
		r.timings.elaspedTime += r.timings.regularInterval

		log.Printf("Round: %d --------------------------------------- Time Elapsed: %s", r.round, r.timings.elaspedTime)
		goldenRound, score := r.determineScore()

		// Select particiapnt and award
		log.Println("Selecting Participant At Random")
		participantIndex := rand.Intn(len(participants))

		log.Printf("Selected %s", participants[participantIndex].Participant)
		participants[participantIndex].Score += score

		scoreBoard := append(participants[:0:0], participants...)
		sort.Slice(scoreBoard, func(i, j int) bool {
			return scoreBoard[i].Score > scoreBoard[j].Score
		})

		// Publish update on websocket
		log.Println("Publishing Update To WebSocket")
		r.window.Send(&raffleUpdate{
			Event: &gotron.Event{Event: "raffle-update"},
			Update: raffleUpdateAttributes{
				Round:               r.round,
				GoldenRound:         goldenRound,
				SelectedParticipant: participants[participantIndex],
				ScoreBoard:          scoreBoard,
			},
		})
	}
}

func (r Raffler) determineScore() (goldenRound, int) {
	goldenRoundStatus := goldenRound{}
	if r.timings.elaspedTime == r.timings.goldenInterval-r.timings.regularInterval {
		log.Printf("Golden Round in %s seconds", r.timings.regularInterval)
		goldenRoundStatus.Next = true
	}
	scoreIncrement := 1
	if r.timings.elaspedTime >= r.timings.goldenInterval {
		log.Println("Golden Round!!!!")
		r.timings.elaspedTime = 0 * time.Second
		scoreIncrement = 5
		goldenRoundStatus.Now = true
	}
	return goldenRoundStatus, scoreIncrement
}
