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

	eion.updateName("Smith")

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

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
