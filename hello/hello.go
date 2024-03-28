package main

import "fmt"

const (
  french  = "French"
	spanish = "Spanish"

	englishPrefix = "Hello "
  frenchPrefix  = "Bonjour "
	spanishPrefix = "Hola "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
		break

	case french:
		prefix = frenchPrefix
		break

	default:
		prefix = englishPrefix
		break

	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
