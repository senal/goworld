package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	sunil := person{firstName: "Sunil", lastName: "Perera"}

	fmt.Println(sunil)

}
