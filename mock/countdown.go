package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

//Sleeper is sleep interface
type Sleeper interface {
	Sleep(time.Duration)
}

type defaultSleeper struct{}

func (sleeper *defaultSleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}

func main() {
	ds := defaultSleeper{}
	Countdown(os.Stdout, &ds)
}

//Countdown to 0
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 5; i > 0; i-- {
		fmt.Fprintf(writer, "%d \n", i)
		sleeper.Sleep(1 * time.Second)
	}
	fmt.Fprintf(writer, "Go!")
}
