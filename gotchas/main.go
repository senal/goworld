package main

import "fmt"

type location struct {
	longitude float64
	altittude float64
}

func main() {

	/*
		* Value Types in Go
			int, float, string, bool and struct

		* Reference types in Go
			slice, maps, channel, pointer, function


	*/

	name := "Bill"
	namePointer := &name
	fmt.Println(&namePointer)
	printPointer(namePointer)

	//mySlice := []string{"Hi", "enjoy", "your", "Day"}
	//updateSlice(mySlice)
	//fmt.Println(mySlice)

}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func updateSlice(slice []string) {
	slice[0] = "Hell0"
}

func changeLatitude(lo location, longitude float64) location {
	lo.longitude = longitude
	return lo
}

func (lo *location) changeLatitude(newLongitude float64) {
	(*lo).longitude = newLongitude
}
