package main

import "fmt"

const spanishPrefix = "Hola "
const frenshPrefix = "Bonjour "

//Hello returns hello
func Hello(name string, lang string) string {
	prefix := "Hello "
	sayName := "World"
	if name != "" {
		sayName = name
	}

	switch lang {
	case "es":
		prefix = spanishPrefix
	case "fr":
		prefix = frenshPrefix
	}

	return prefix + sayName
}

//SayHello print hello to the screen
func SayHello() {
	fmt.Println(Hello("Eslam", "en"))
}

func main() {
	SayHello()
}
