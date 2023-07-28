package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Action",
		contact: contact{
			email:   "r@r.com.au",
			zipCode: 1234,
		},
	}

	jimmPointer := &jim
	jimmPointer.updateFirstName("Jimmy")
	jim.print()
}
