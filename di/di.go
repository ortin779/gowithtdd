package di

import "io"

const (
	helloEnglishPrefix = "Hello, "
)

func Greet(w io.Writer, name string) {
	message := helloEnglishPrefix + name
	w.Write([]byte(message))
}
