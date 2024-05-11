package main

import (
	"fmt"
	"io"
  "os"
  "time"
)

type Sleeper interface {
  Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
  sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(buffer, "Go")
}
