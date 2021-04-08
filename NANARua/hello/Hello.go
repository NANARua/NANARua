package main

import "fmt"

const englishHelloPrefix = "Hello "
const frenchHelloPrefix = "Bonjour "
const spanishHelloPrefix = "Hola "

func Hello(name string, language string) string {
	if name == "" {
		name = "world!"
	}
	prefix := englishHelloPrefix
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("yiqi", "Spanish"))
}
