package main

import (
	"os"
	"time"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep})
}
