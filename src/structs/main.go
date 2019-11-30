package main

import "fmt"

func main() {
	/*
		// one way of initializing struct
		alex := person{"Alex", "Jandon"}
		fmt.Println(alex)

		// another way of initialization struct
		james := person{firstName: "James", lastName: "Anderson"}
		fmt.Println(james)

		//third way of initialization
		var jimmmy person
		// by default it will take default vaues - for sting its "", for int and float its 0, for bool its false
		fmt.Println(jimmmy)
		jimmmy.firstName = "Joe"
		jimmmy.lastName = "Root"

		fmt.Printf("%+v", jimmmy)
	*/

	eion := person{
		firstName: "Eion",
		lastName:  "Morgan",
		contactInfo: contactInfo{
			email:   "admin@gamil.com",
			zipcode: 500081,
		},
	}

	// fmt.Printf("%+v", eion)
	eion.printPerson()

	personPointer := &eion // &variable -> gives me the memory address of the value is pointing at.

	personPointer.updateName("Smith")

	eion.printPerson()
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

type person struct {
	firstName   string
	lastName    string
	contactInfo // equivalent to contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}

func (pointer *person) updateName(newFirstName string) { // *persion -> this is a type description - that means we are working with a pointer to a person.
	(*pointer).firstName = newFirstName // *pointer -> give me the value this memory address is pointing at
}
