package mixer

import "sync"

type Source interface {
	Write([]float64)
}

type Track struct {
	Source Source
	Volume float64
}

type Mixer struct {
	Tracks []Track
	lock   sync.Mutex
}

func NewMixer() *Mixer {
	return &Mixer{}
}

func (m *Mixer) Write(buf []float64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	trackData := make([][]float64, len(m.Tracks))
	totalVol := 0.0
	for i, t := range m.Tracks {
		trackData[i] = make([]float64, len(buf))
		t.Source.Write(trackData[i])
		totalVol += t.Volume
	}

	for i := 0; i < len(buf); i++ {
		if totalVol == 0 {
			buf[i] = 0
			continue
		}

		tot := 0.0
		for j, t := range trackData {
			tot += t[i] * m.Tracks[j].Volume
		}
		buf[i] = tot / totalVol
	}
}

func (m *Mixer) AddTrack(source Source, volume float64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.Tracks = append(m.Tracks, Track{
		Source: source,
		Volume: volume,
	})
}
