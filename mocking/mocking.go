package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, err := fmt.Fprintln(out, i)
		if err != nil {
			log.Fatalf("Error printing countdown: %v", err)
		}
		sleeper.Sleep()
	}
	_, err := fmt.Print(out, finalWord)
	if err != nil {
		log.Fatalf("Error printing countdown: %v", err)
	}
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
