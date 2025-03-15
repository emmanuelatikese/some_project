package other_funcs

import(
	"fmt"
    "io"
    "os"
    "github.com/hajimehoshi/go-mp3"
    "github.com/go-audio/audio"
    "github.com/go-audio/wav"
)

func ConvertMP3ToWav(filename string) (string, error){

    mp3File, err := os.Open(filename);
    if err != nil {
        return "", fmt.Errorf("error: %v", err);
        
    }

    defer mp3File.Close();

    mp3Decoder, err := mp3.NewDecoder(mp3File);
    if err != nil {
        return "", fmt.Errorf("error: %v", err);
    }


    wavFile, err := os.Create("output.wav");
    if err != nil{
        return "", fmt.Errorf("error: %v", err);
    }

    wavEncoder := wav.NewEncoder(wavFile,
            44100,           // Sample rate
            16,              // Bits per sample
            2,               // Number of channels
            1)              // WAV format (1 = PCM)

    buf := make([]byte, 1024);

    audioBuf := &audio.IntBuffer{
        Format: &audio.Format{
            NumChannels: 2,
            SampleRate: 44100,
        },
    }

    for {
        n, err := mp3Decoder.Read(buf);
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", fmt.Errorf("error: %v", err);
        }

        fmt.Println(n);

        samples := make([]int, n/2)
        for i := 0; i < n/2; i += 2 {
            samples[i] = int(int16(buf[i]) | int16(buf[i+1])<<8)
        }
        
        audioBuf.Data = samples;
        err = wavEncoder.Write(audioBuf);
        if err != nil {
            return "", fmt.Errorf("error: %v", err);
        }
    }

	return "output.wav", nil;
}