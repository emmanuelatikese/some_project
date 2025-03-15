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

    sampleRate := mp3Decoder.SampleRate();


    wavFile, err := os.Create("output.wav");
    if err != nil{
        return "", fmt.Errorf("error: %v", err);
    }

    wavEncoder := wav.NewEncoder(wavFile,
            sampleRate,           // Sample rate
            16,              // Bits per sample
            2,               // Number of channels
            1)              // WAV format (1 = PCM)

    buf := make([]byte, 8192);

    audioBuf := &audio.IntBuffer{
        Format: &audio.Format{
            NumChannels: 2,
            SampleRate: sampleRate,
        },
    }

    for {
        n, err := mp3Decoder.Read(buf);
        if err != nil {
            if err == io.EOF || n == 0 {
                break
            }
            return "", fmt.Errorf("error: %v", err);
        }


        samples := make([]int, n/2)
        for i := 0; i < n-1; i += 2 {
                sample := int(int16(buf[i]) | int16(buf[i+1])<<8)
                samples[i/2] = sample
        }
        
        audioBuf.Data = samples;
        err = wavEncoder.Write(audioBuf);
        if err != nil {
            return "", fmt.Errorf("error: %v", err);
        }
    }

    if err := wavEncoder.Close(); err != nil {
        return "", fmt.Errorf("error: %v", err);
    }

	return "output.wav", nil;
}