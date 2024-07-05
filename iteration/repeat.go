package iteration

const (
	repeatCount = 5
)

func Repeat(char rune) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += string(char)
	}
	return repeated
}
