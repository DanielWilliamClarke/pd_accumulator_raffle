package accumulator

import (
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/Equanox/gotron"
)

type PartiticpantScore struct {
	Participant string `json:"participant"`
	Score       int    `json:"score"`
}

func ParseParticiapants(participants []string) (ps []PartiticpantScore) {
	for _, p := range participants {
		ps = append(ps, PartiticpantScore{
			Participant: p,
			Score:       0,
		})
	}
	return ps
}

type RaffleUpdateAttributes struct {
	Round               int                 `json:"round"`
	SelectedParticipant PartiticpantScore   `json:"selectedParticipant"`
	ScoreBoard          []PartiticpantScore `json:"scoreBoard"`
}

type RaffleUpdate struct {
	*gotron.Event
	Update RaffleUpdateAttributes `json:"update"`
}

type Raffler struct {
	Participants []PartiticpantScore
	Window       *gotron.BrowserWindow
}

func (r Raffler) Run() {

	log.Println("Starting Accumulator Raffle")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	round := 0

	for {

		time.Sleep(20 * time.Second)

		round++

		log.Println("Selecting Participant At Random")
		participantIndex := r1.Intn(len(r.Participants))

		log.Printf("Selected %s", r.Participants[participantIndex].Participant)
		r.Participants[participantIndex].Score++

		scoreBoard := append(r.Participants[:0:0], r.Participants...)
		sort.SliceStable(scoreBoard, func(i, j int) bool {
			return scoreBoard[i].Score < scoreBoard[j].Score
		})

		log.Println("Publishing Update To WebSocket")
		r.Window.Send(&RaffleUpdate{
			Event: &gotron.Event{Event: "raffle-update"},
			Update: RaffleUpdateAttributes{
				Round:               round,
				SelectedParticipant: r.Participants[participantIndex],
				ScoreBoard:          scoreBoard,
			},
		})
	}
}
