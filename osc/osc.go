package osc

import (
	"sync"
)

type OSC struct {
	freq      float64
	rate      float64
	amplitude float64
	generator Generator
	lock      sync.RWMutex
}

type Generator interface {
	Next(freq, rate float64) float64
}

func NewOSC(freq, rate float64, generator Generator) *OSC {
	return &OSC{
		freq:      freq,
		rate:      rate,
		amplitude: 1.0,
		lock:      sync.RWMutex{},
		generator: generator,
	}
}

func (o *OSC) Write(buffer []float64) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	for i := 0; i < len(buffer); i++ {
		buffer[i] = o.generator.Next(o.freq, o.rate) * o.amplitude
	}
}

func (o *OSC) Freq() float64 {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return o.freq
}

func (o *OSC) Amp() float64 {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return o.amplitude
}

func (o *OSC) Rate() float64 {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return o.rate
}

func (o *OSC) SetFreq(f float64) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.freq = f
}

func (o *OSC) SetAmp(a float64) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	o.amplitude = a
}
