package utils

import (
	"math"
)

var R = math.Pow(2.0, 1.0/12.0)

type Scale struct {
	BaseFreq float64
	Note     int
	Scale    int
}

const (
	AFlat  = -1
	A      = 0
	ASharp = 1
	BFlat  = 1
	B      = 2
	C      = -9
	CSharp = -8
	DFlat  = -8
	D      = -7
	DSharp = -6
	EFlat  = -6
	E      = -5
	F      = -4
	FSharp = -3
	GFlat  = -3
	G      = -2
	GSharp = -1
)

func (s Scale) ToFreq(note int, scale int) float64 {
	noteDiff := note - s.Note
	res := s.freqForScale(scale)
	if noteDiff == 0 {
		return res
	}

	if noteDiff < 0 {
		for i := 0; i > noteDiff; i-- {
			res = res / R
		}
	}
	if noteDiff > 0 {
		for i := 0; i < noteDiff; i++ {
			res = res * R
		}
	}

	return res
}

func (s Scale) freqForScale(scale int) float64 {
	scaleDiff := s.Scale - scale
	return s.BaseFreq / math.Pow(2, float64(scaleDiff))
}

func NoteToString(note int) string {
	switch note {
	case AFlat:
		return "Ab"
	case A:
		return "A"
	case ASharp:
		return "A#"
	case B:
		return "B"
	case C:
		return "C"
	case CSharp:
		return "C#"
	case D:
		return "D"
	case DSharp:
		return "D#"
	case E:
		return "E"
	case F:
		return "F"
	case FSharp:
		return "F#"
	case G:
		return "G"
	}
	return "n/c"
}
