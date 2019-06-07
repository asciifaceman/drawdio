package intake

import (
	"azul3d.org/engine/audio"
)

// Sample is a de-interleaved raw audio sample second
type Sample struct {
	Timecode int
	S        audio.Float64
}

// SampleSet represents many samples
type SampleSet struct {
	Samples []*Sample
}

// Decompose returns just a slice of all float64s
func (s *SampleSet) Decompose() []float64 {
	var samp []float64
	for _, part := range s.Samples {
		for _, sample := range part.S {
			samp = append(samp, sample)
		}
	}
	return samp
}

/*
// Sample is a de-interleaved raw audio sample second
type Sample struct {
	Timecode int
	L        audio.Float64
	R        audio.Float64
}

// SampleSet represents many samples
type SampleSet struct {
	Samples []*Sample
}

// DeInterleave splits L/R channels
func DeInterleave(input audio.Float64) *Sample {
	var liter, riter int
	l := make(audio.Float64, input.Len()/2)
	r := make(audio.Float64, input.Len()/2)

	for iter, sample := range input {
		if iter%2 == 0 {
			l.Set(liter, sample)
			liter++
		} else {
			l.Set(riter, sample)
			riter++
		}
	}

	return &Sample{
		L: l,
		R: r,
	}
}
*/
