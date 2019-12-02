package main

import "fmt"

/**

is line of code right here can be interpreted as essentially saying hey everyone inside of this

program everyone inside the program all of you different types out there all of you different types

like English spot in Spanish box to all of you to whom it may concern our program now has a new type

a third custom type called bot.

So we have a new custom type called bot right here.

Now there's nothing special about the named but I only use the named bot to kind of you know kind of

implement or kind of give you the idea of a more kind of generic approach to an English bot or Spanish

spot.

So we're you can kind of imagine we are defining some behaviors right here that are common to all different

bots that we might have.

So if we also had a Russian bot a on an old French box Italian bot and so on they all kind of share

the same implementation from this type based bot right here.

So to all these custom types just know to whom it may concern.

There's a new custom type called box and then inside the definition of this custom type we said if you

are a type in this program with a function called get a greeting and you return a type of string then

you are now also an honorary member of type bot.
*/

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("Welcome to go interfaces")
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
