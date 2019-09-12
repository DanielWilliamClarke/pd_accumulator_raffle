package accumulator

import (
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/Equanox/gotron"
)

type RaffleUpdate struct {
	*gotron.Event
	Update RaffleUpdateAttributes `json:"update"`
}

type RaffleUpdateAttributes struct {
	Round               int                 `json:"round"`
	GoldenRound         GoldenRound         `json:"goldenRound"`
	SelectedParticipant PartiticpantScore   `json:"selectedParticipant"`
	ScoreBoard          []PartiticpantScore `json:"scoreBoard"`
}

type GoldenRound struct {
	Next bool `json:"next"`
	Now  bool `json:"now"`
}

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

type Raffler struct {
	Participants []PartiticpantScore
	Window       *gotron.BrowserWindow
}

func (r Raffler) Run() {

	log.Println("Starting Accumulator Raffle")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	round := 0

	elaspedTime := time.Duration(0)
	regularInterval := 20 * time.Second
	goldenInterval := 300 * time.Second

	for {

		time.Sleep(regularInterval)

		round++
		elaspedTime += regularInterval
		log.Printf("Round: %d --------------------------------------- Time Elapsed: %s", round, elaspedTime)

		goldenRoundNext := false
		if elaspedTime == goldenInterval-regularInterval {
			log.Printf("Golden Round in %s seconds", regularInterval)
			goldenRoundNext = true
		}
		scoreIncrement := 1
		isGoldenRound := false
		if elaspedTime == goldenInterval {
			scoreIncrement = 5
			elaspedTime = 0
			isGoldenRound = true
		}

		log.Println("Selecting Participant At Random")
		participantIndex := r1.Intn(len(r.Participants))

		log.Printf("Selected %s", r.Participants[participantIndex].Participant)
		r.Participants[participantIndex].Score += scoreIncrement

		scoreBoard := append(r.Participants[:0:0], r.Participants...)
		sort.Slice(scoreBoard, func(i, j int) bool {
			return scoreBoard[i].Score > scoreBoard[j].Score
		})

		log.Println("Publishing Update To WebSocket")
		r.Window.Send(&RaffleUpdate{
			Event: &gotron.Event{Event: "raffle-update"},
			Update: RaffleUpdateAttributes{
				Round: round,
				GoldenRound: GoldenRound{
					Next: goldenRoundNext,
					Now:  isGoldenRound,
				},
				SelectedParticipant: r.Participants[participantIndex],
				ScoreBoard:          scoreBoard,
			},
		})
	}
}
