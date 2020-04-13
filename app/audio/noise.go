// go get -u -v -f all
// go build audio.go
// go build V:\Programming\JavaScript\Programming\Projects\azatom-keep-alive\app\audio\audio.go

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("hello, world\n")
// }

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
)

// Noise is good
func Noise() beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			fmt.Println(i)
			samples[i][0] = 0.25
			samples[i][1] = 0.25
			// samples[i][0] = rand.Float64()*2 - 1
			// samples[i][1] = rand.Float64()*2 - 1
		}
		return len(samples), true
	})
}

func main() {

	volumePtr := flag.Float64("volume", 5, "Audio volume (0 is normal volume, <0 quieter, >0 louder)")
	delayPtr := flag.Float64("delay", 3, "Time between audio repeats (seconds)")
	repeatPtr := flag.Int("repeats", 3, "How many times the audio will play. 0 = infinity")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	// fmt.Println("executing go 1", rand.Float64()*2-1)

	sr := beep.SampleRate(200)
	// speaker.Init(sr, sr.N(time.Second/100))
	speaker.Init(sr, int(sr))
	audiobuffer := Noise()
	// loop := &beep.Ctrl{Streamer: beep.Loop(-1, audiobuffer), Paused: false}
	// audiobuffer := randomDataStreamer()
	volume := &effects.Volume{
		Streamer: audiobuffer,
		Base:     2,
		Volume:   *volumePtr,
		Silent:   false,
	}
	// done := make(chan bool)

	for k := 1; k <= *repeatPtr || *repeatPtr == 0; k++ {
		speaker.Play(beep.Seq(beep.Take(sr.N(200*time.Millisecond), volume), beep.Callback(func() {
			// done <- true
		})))
		// // select {}
		// <-done
		// speaker.Play(beep.Take(5, volume))
		// speaker.Play(volume)

		time.Sleep(time.Millisecond * 1) // line required for the sound to get played
		speaker.Clear()
		fmt.Printf("Waiting for %v seconds\n", *delayPtr)
		time.Sleep(time.Second * time.Duration(*delayPtr))
	}
}
