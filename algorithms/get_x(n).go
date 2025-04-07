package algo

import (
	"fmt"
	"os"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)


func GetXnFromWav(filename *os.File) ([]float64, error) {

	wavDecoder := wav.NewDecoder(filename)
	if !wavDecoder.IsValidFile() {
		return nil, fmt.Errorf("invalid WAV file: %v", wavDecoder.Err())
	}

	format := wavDecoder.Format()
	numChannels := int(format.NumChannels)

	const bufferSize = 8192
	buf := &audio.IntBuffer{
		Format: format,
		Data:   make([]int, bufferSize*numChannels),
	}


	var samples []float64

	for {

		n, err := wavDecoder.PCMBuffer(buf)
		if err != nil {
			break 
		}
		fmt.Println("hey")

		for i := 0; i < n*numChannels; i++ {
			sample := float64(buf.Data[i])

			switch wavDecoder.BitDepth {
			case 8:
				sample = (sample - 128) / 128.0
			case 16:
				sample = sample / 32768.0
			case 24:
				sample = sample / 8388608.0
			case 32:
				sample = sample / 2147483648.0
			default:
				return nil, fmt.Errorf("error on bitDepth: %d", wavDecoder.BitDepth)
			}

			samples = append(samples, sample)
		}
	}

	return samples, nil
}
