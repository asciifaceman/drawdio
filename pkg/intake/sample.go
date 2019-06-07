package intake

import (
	"fmt"

	"azul3d.org/engine/audio"
	// WAV
	_ "azul3d.org/engine/audio/wav"
	"go.uber.org/zap"
)

// Sample collects the contents of the wav
// file and returns the sample collection or error
func (i *Intake) Sample() (*SampleSet, error) {

	i.Logger.Info("Beginning sampling run...",
		zap.String("filename", i.c.Filename),
		zap.Uint("resolution", i.resolution),
	)

	// create the decoder
	decoder, _, err := audio.NewDecoder(i.r)
	if err != nil {
		return nil, err
	}

	config := decoder.Config()

	seconds := 1
	sampleSet := &SampleSet{}
	for {
		samples := make(audio.Float64, uint(config.SampleRate*config.Channels)/i.resolution)
		read, err := decoder.Read(samples)
		if err != nil && err != audio.EOS {
			return nil, err
		}
		if err == audio.EOS {
			break
		}

		s := &Sample{
			Timecode: seconds,
			S:        samples,
		}
		sampleSet.Samples = append(sampleSet.Samples, s)

		i.Logger.Info("Read samples...",
			zap.Int("count", read),
			zap.Int("seconds", seconds),
		)

		seconds++
	}

	//spew.Dump(sampleSet)
	fmt.Printf("Found %d sample sets\n", len(sampleSet.Samples))

	return sampleSet, nil
}
