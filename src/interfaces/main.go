package main

import "fmt"

/**

This line below can be interpreted as essentially saying

hey everyone inside of this program everyone inside the program all of you different types out there all of you different types
like English bot in Spanish bot to all of you to whom it may concern our program now has a new type a third custom type called bot.
*/

/**

So we have a new custom type called bot right here.

So we're you can kind of imagine we are defining some behaviors right here that are common to all different bots that we might have.

So if we also had a Russian bot a on an old French box Italian bot and so on they all kind of share the same implementation from this type based bot right here.

So to all these custom types just know to whom it may concern. There's a new custom type called box and then inside the definition of this custom type we said if you

are a type in this program with a function called get a greeting and you return a type of string then you are now also an honorary member of type bot.
*/

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}
type russianBot struct{}

func main() {
	fmt.Println("Welcome to go interfaces")
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// rb := russianBot{}
	// printGreeting(rb) will not work here untill unless I m not going to implement getGreeting func from interface with russian bot
}

// func (rb russianBot) getGreeting() string {
// 	return "Hi From Russia!!!!"
// }

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

/**

There's just a couple of more quick notes I want to give you an interface's and then we're going to

write up a couple more quick examples with interface's to just get a better sense of how they work.

So let's go through these last couple of months.

All right so first off interfaces are not generic types so if you're coming from say C-Sharp or Java

which have access to generic types interfaces are not a full on replacement for them.

So we can't use interfaces to write completely generic code or code that accepts absolutely any kind

of type.

And then just kind of deals with it appropriately go actually somewhat famously does not have support

for generic types.

And so it's best not to think of interfaces as being a replacement for them.

They're just a tool that we have to sometimes write code that is a little bit more concise or doesn't

repeat itself quite as much as it might otherwise.

Now if you're not familiar with generic types that's really totally fine don't worry about it.

I just want to throw that in there as a very quick aside.

Now the next thing I want to tell you is that interfaces are what we call implicit.

So inside of the our code snippet that we put together over here you'll notice that at no point in time

did we ever have to say oh English not in Spanish spot just so you know.

You are also of type but because you satisfy this interface all we did was to clear this interface and

the different functions that it expects to see and then go essentially took it from there.

So you and I never put together any type of manual code to somehow linked together bought an English

pot in Spanish sponte go takes care of all the matching for us and so because God takes care of it for

us we say that interfaces are implicit.

So we did not have to explicitly define any link between these different types here.

Now that's both a blessing and a curse.

Personally I think that it's really nice that we don't have to write out some more boilerplate to say

something like type English bought implements bought or something like that.

You know we don't have to do anything like that in go which is kind of nice.

But on the flip side it makes figuring out what types of yours implement other interfaces.

A little bit challenging because you kind of have to flip between the interface definition and your

custom type and say OK I need to make it greeting and it has to return a string and then you go down

here and you start working on your function and then you say oh oops I returned an into needs to be

a string and you don't necessarily get a lot of quick feedback on exactly how to implement that interface

until you start to see some error messages pop up when you compile your code.

So this whole interfaces implicit thing kind of both a blessing and a curse because it saves us from

some boilerplate but also sometimes makes it a little bit hard to keep track of which types implement

which interfaces.

Now next point interfaces are a contract to help us manage types.

And this is an extremely important note something that I really want to point out here.

I've seen a lot of people look at interfaces and they think oh interfaces are like some type of built

in test for my code or they're just going to make sure that my code always works correctly.

And you really do not want to think of interfaces in such fashion.

You're on the right hand side.

Notice the note I had it on I said.

Garbage in equals garbage out.

So if you are implementing some function to match get greeting right here and you return some you know

maybe return a value of type string but it's like just a garbage value or a garbage string.

Hey goes going to think that you're still implementing greeting right here but your program might crash

later on or might like not do exactly what you expect.

In other words let's look at a very practical example of this.

So when we define get greeting for both Spanish and English spot right here you and I know very well

that OK like this needs to return a type of greeting.

But if we instead decide to return something like this right here.

Well hey it's still of type string.

It doesn't really probably work for what we're trying to do.

Like we want to return a greeting something like Hi there or hello or how are you doing from this function

but all God cares about is making sure that your teacher turned the correct type.

So interfaces are about helping you reuse code or kind of form a contract between different functions

and different types and they're not necessarily going to make sure that you don't make some otherwise

basic logic mistake inside of your code.

And we'll see some really good examples of that when we start looking at interfaces inside of the standard

library I'm going to point out some examples and we're raving about like write out an example or two

where yeah hey interface is going to help us but it's not necessarily going to keep us from making some

big mistake in our actual logic inside of our codebase.

OK and now last note here really important note for you as someone who is learning go remember interfaces

are tough tough really tough.

Now step number one in figuring out how interfaces work and what they're doing for us is just understanding

what interfaces do and how to read the official standard library documentation and understand the different

interfaces that are defined in there.

So just understanding just like kind of reading interfaces is step number one in this difficulty curve.

So do not worry from the get go about writing your own interfaces and making use of them inside of your

own application.

Like I'll tell you right now you can write entire just lines and thousands of lines of code with us

go code in your own custom program in your own personal projects without worrying about writing your

own interfaces so interfaces are not a requirement of the language you don't have to create your own.

Of course it's highly recommended that you do at some points for code quality and clarity and all that

kind of good stuff.

But you are not necessarily required to use them or to write your own interfaces.

So step one and the difficulty curve like Step 1 in just getting used to interfaces is being able to

understand the standard library documentation and understand what the standard library documentation

means when it says oh yeah I expect something of type blah blah blah and that type ends up being an

interface and so that's what we're going to focus on here for number one is just how to look at the

standard library documentation and understand the use of interfaces in there now much later on.

You know as you start to gather more experience around interfaces well then at some point in the future.

OK then at that point once you feel more established once you've got a lot more practice then come back

and start thinking about hey how am I going to write my own custom interfaces for my own application.

*/
