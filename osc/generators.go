package osc

import "math"

func NewSinGenerator() *sinGenerator {
	return &sinGenerator{}
}

func NewSynthGenerator() *synthGenerator {
	return &synthGenerator{
		f1: NewSinGenerator(),
		f2: NewSinGenerator(),
		f3: NewSinGenerator(),
	}
}

type sinGenerator struct {
	phase float64
}

func (s *sinGenerator) Next(freq, rate float64) float64 {
	phaseShift := 2.0 * math.Pi * freq / rate
	s.phase += phaseShift
	if s.phase > 2.0*math.Pi {
		s.phase -= 2.0 * math.Pi
	}
	return math.Sin(s.phase)
}

type synthGenerator struct {
	f1 *sinGenerator
	f2 *sinGenerator
	f3 *sinGenerator
}

func (s *synthGenerator) Next(freq, rate float64) float64 {
	return 0.1*s.f1.Next(freq/4, rate) +
		0.6*s.f2.Next(freq, rate) +
		0.2*s.f3.Next(freq*2, rate)
}
