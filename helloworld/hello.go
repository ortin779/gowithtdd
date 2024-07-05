package helloworld

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	teleguHelloPrefix  = "Namasthey, "

	spanish = "Spanish"
	french  = "French"
	telugu  = "Telugu"

	defaultName = "world"
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case telugu:
		prefix = teleguHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Nitro", "English"))
}
