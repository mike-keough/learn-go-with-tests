package main

const (
	spanish       string = "Spanish"
	french        string = "French"
	frenchPrefix  string = "Bonjour, "
	englishPrefix string = "Hello, "
	spanishPrefix string = "Hola, "
)

func main() {
	Hello("Mike", "Spanish")
}

func Hello(n, l string) string {
	if n == "" {
		n = "World"
	}

	return greetingPrefix(l) + n
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}
