package main

import "fmt"

type bot interface {
	// englishBot and spanishBot both have receiver functions called "getGreeting"
	// the interface is saying that any type that has getGreeting defined as a
	// receiver function a member of the "bot" type
	// now printGreeting below can call getGreeting for any type that his it defined
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}


func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// can omit value "eb" because its not used. interesting!
func (englishBot) getGreeting() string {
	return "Hello There"
}

func (sb spanishBot) getGreeting() string {
	return "Hola Senior"
}


