package main

import "fmt"

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// can omit value "eb" because its not used. interesting!
func (englishBot) getGreeting() string {
	return "Hello There"
}

func (sb spanishBot) getGreeting() string {
	return "Hola Senior"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}