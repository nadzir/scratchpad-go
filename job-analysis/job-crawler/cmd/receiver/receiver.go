package main

import (
	"fmt"

	queue "github.com/nadzir/scratchpad-go/job-analysis/pkg/queue/receive"
)

func main() {
	channel := make(chan []byte)
	go queue.StartReceiver(channel)

	for {
		select {
		case msg1 := <-channel:
			fmt.Println("received", string(msg1))
		}
	}
}
