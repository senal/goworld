package main

import "fmt"

func main() {
	// ways to make an emtpy map
	//1. var colors map[string]string
	colors := make(map[string]string)

	// how to append elements to an empty map
	colors["red"] = "this is red"
	colors["white"] = "this is white"

	// colors := map[string]string{
	// 	"red":   "this is red",
	// 	"green": "this is green",
	// }

	// delete elements of a map
	delete(colors, "red")

	fmt.Println(colors)
}
