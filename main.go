package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gordonklaus/portaudio"
	mixer "github.com/johnsudaar/sound/mixers"
	"github.com/johnsudaar/sound/osc"
	"github.com/johnsudaar/sound/utils"
)

func main() {
	bufferSize := 64
	rate := 48000.0

	portaudio.Initialize()
	defer portaudio.Terminate()

	m := mixer.NewMixer()

	out := make([]float32, bufferSize)
	buf := make([]float64, bufferSize)
	stream, err := portaudio.OpenDefaultStream(0, 1, rate, bufferSize, &out)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		panic(err)
	}
	defer stream.Stop()
	go noteGenerator(rate, m, true)

	for {
		m.Write(buf)
		utils.F64ToF32(buf, out)
		err := stream.Write()
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func noteGenerator(rate float64, m *mixer.Mixer, first bool) {
	s := utils.Scale{
		BaseFreq: 440.0,
		Note:     utils.A,
		Scale:    4,
	}
	oscillator := osc.NewOSC(1, rate, osc.NewSynthGenerator())
	if first {
		m.AddTrack(oscillator, 3)
	} else {
		m.AddTrack(oscillator, 1)
	}

	notes := []int{
		62, 64, 66, 67, 69, 71, 73, 74, // D Maj ascending
		57, 59, 61, 62, 64, 66, 67, 69, // A Maj ascending
		59, 61, 62, 64, 66, 67, 69, 71, // B Min ascending
		54, 56, 57, 59, 61, 62, 64, 66, // F# Min ascending
		55, 57, 59, 60, 62, 64, 66, 67, // G Maj ascending
		62, 64, 66, 67, 69, 71, 73, 74, // D Maj ascending
		55, 57, 59, 61, 62, 64, 66, 67, // G Maj ascending
		57, 59, 61, 62, 64, 66, 67, 69, // A Maj ascending
	}

	firstRun := true
	for {
		for i, n := range notes {
			if first && firstRun && (i == 16 || i == 32 || i == 48) {
				go noteGenerator(rate, m, false)
			}
			scale := n/12 - 0
			note := n%12 - 9
			if first {
				scale -= 3
			}
			freq := s.ToFreq(note, scale)
			oscillator.SetFreq(freq)
			fmt.Printf("Playing: %v\t%s\n", scale, utils.NoteToString(note))
			time.Sleep(500 * time.Millisecond)
		}
		firstRun = false
	}

}
