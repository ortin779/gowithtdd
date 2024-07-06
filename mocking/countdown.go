package main

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}
