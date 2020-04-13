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
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// args := os.Args[1:]
	volumePtr := flag.Float64("volume", 0, "Audio volume (0 is normal volume, <0 quieter, >0 louder)")
	delayPtr := flag.Float64("delay", 600, "Time between audio repeats (seconds)")
	repeatPtr := flag.Uint("repeats", 0, "How many times the audio will play. 0 = infinity")

	wordPtr := flag.String("string", "foo", "a string")
	numbPtr := flag.Int("integer", 2, "an integer")
	floatPtr := flag.Float64("float64", 1.23456, "a decimal")
	boolPtr := flag.Bool("boolean", false, "a boolean")
	flag.Parse()

	fmt.Println("string:", *wordPtr)
	fmt.Println("integer:", *numbPtr)
	fmt.Println("float:", *floatPtr)
	fmt.Println("boolean:", *boolPtr)
	fmt.Println("args:", flag.Args())
	// f, err := os.Open(".\\click.mp3")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// streamer, format, err := mp3.Decode(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer streamer.Close()

	// speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// done := make(chan bool)
	// speaker.Play(beep.Seq(streamer, beep.Callback(func() {
	// 	done <- true
	// })))

	// <-done

	f, err := os.Open(".\\click.mp3")
	// f, err := ioutil.ReadFile(".\\click.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	// streamer.Close()
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	k := uint(1)
	for k <= *repeatPtr || *repeatPtr == 0 {
		// speaker.Play(shot)

		audio := buffer.Streamer(0, buffer.Len())

		ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, audio), Paused: false}
		volume := &effects.Volume{
			Streamer: ctrl,
			Base:     2,
			Volume:   float64(*volumePtr),
			Silent:   false,
		}
		speaker.Play(volume)

		// speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		// speaker.Play(beep.Seq(audio, beep.Callback(func() {
		// 	// done <- true
		// })))
		fmt.Println("Loop", k, flag.Args())
		time.Sleep(time.Second * 1)
		speaker.Clear()
		time.Sleep(time.Second * time.Duration(*delayPtr))
		k++
	}

}
