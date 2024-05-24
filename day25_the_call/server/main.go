package main

import (
	"fmt"
	"log"
	in "server/input"
	"server/output"
)

func main() {
	input, err := in.New(8192)
	if err != nil {
		log.Fatal(err)
	}
	output, err := output.New(8192, input.GetStream())
	if err != nil {
		log.Fatal(err)
	}
	stop := make(chan struct{})
	go func() {
		for {
			var s string
			fmt.Print("Press a command (play, pause, stop): ")
			fmt.Scanln(&s)
			switch s {
			case "pause":
				fmt.Println("Paused")
				input.Pause()
			case "play":
				fmt.Println("Continue")
				input.Play()
				output.Play()
			case "stop":
				fmt.Println("Prepare to stop all....")
				input.Stop()
				output.Stop()
				stop <- struct{}{}
				return
			default:
				fmt.Println("Wrong command!!!")
				continue

			}
		}
	}()
	<-stop
	// signal := make(chan struct{})
	// go func() {
	// 	for {
	// 		err := binary.Read(input.GetStream(), binary.LittleEndian, &buf)
	// 		reading.Write()
	// 		if err != nil {
	// 			signal <- struct{}{}
	// 			return
	// 		}
	// 	}
	// }()

	// <-signal
}

//caller and callee!!!
