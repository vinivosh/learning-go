package main

import "fmt"

const (
	english               = "en"
	spanish               = "es"
	french                = "fr"
	portuguese            = "pt"
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

func getGreetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return getGreetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("world", english))
}
