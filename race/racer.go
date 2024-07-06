package race

import (
	"fmt"
	"net/http"
)

func Racer(a, b string) string {
	var winner string

	select {
	case <-ping(a):
		fmt.Println(a)
		winner = a
	case <-ping(b):
		fmt.Println(b)
		winner = b
	}
	return winner
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
