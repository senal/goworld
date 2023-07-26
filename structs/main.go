package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Action",
		contact: contactInfo{
			email:   "r@r.com.au",
			zipCode: 1234,
		},
	}

	fmt.Printf("Jim %+v", jim)
}
