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

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Action",
		contact: contact{
			email:   "j@j.com.au",
			zipCode: 1234,
		},
	}

	ranga := person{
		firstName: "Ranga",
		lastName:  "Fernando",
		contact: contact{
			email:   "r@r.com.au",
			zipCode: 1234,
		},
	}

	ranga.updateFirstName("Suranga")
	//jim.updateFirstName("Jimmy")
	jim.print()
	ranga.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(pointerToPerson).firstName = newFirstName
}
