package algo

import (
	"fmt"
	"io"
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
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("error reading PCM buffer: %v", err)
		}
		if n == 0 || err == io.EOF {
			break // No more data to read
		}

		// Process the buffer
		for i := 0; i < n*numChannels && i < len(buf.Data); i++ {
			sample := float64(buf.Data[i])

			switch wavDecoder.BitDepth {
			case 8:
				sample = (sample - 128) / 128.0 // 8-bit is typically unsigned
			case 16:
				sample = sample / 32768.0 // 16-bit signed
			case 24:
				sample = sample / 8388608.0 // 24-bit signed
			case 32:
				sample = sample / 2147483648.0 // 32-bit signed
			default:
				return nil, fmt.Errorf("unsupported bit depth: %d", wavDecoder.BitDepth)
			}
			samples = append(samples, sample)
		}

	}

	if len(samples) == 0 {
		return nil, fmt.Errorf("no samples extracted from WAV file")
	}

	return samples, nil
}