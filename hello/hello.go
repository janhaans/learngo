package main

import "fmt"

const englishHello = "Hello, "
const dutchHello = "Hallo, "
const frenchHello = "Bonjour, "
const italyHello = "Ciao"

var prefix string

// Hello says hello to the world
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	switch language {
	case "dutch":
		prefix = dutchHello
	case "french":
		prefix = frenchHello
	case "italy":
		prefix = italyHello
	default:
		prefix = englishHello
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
