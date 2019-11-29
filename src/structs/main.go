package main

import "fmt"

func main() {

	// one way of initializing struct
	alex := person{"Alex", "Jandon"}
	fmt.Println(alex)

	// another way of initialization struct
	james := person{firstName: "James", lastName: "Anderson"}
	fmt.Println(james)

}

type person struct {
	firstName string
	lastName  string
}
