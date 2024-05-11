package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(buffer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(buffer, "Go")
}
