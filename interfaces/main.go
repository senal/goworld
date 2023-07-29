package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	//sb := spanishBot{}

	printGreeting(eb)
	//printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Holla!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sp spanishBot) {
// 	fmt.Println(sp.getGreeting())
// }
