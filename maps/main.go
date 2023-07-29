package main

import "fmt"

func main() {
	// ways to make an emtpy map
	//1. var colors map[string]string
	colors := map[string]string{
		"red":   "this is  red",
		"white": "this is white",
		"green": "this is green",
		"blue":  "this is blue",
	}
	printElements(colors)

}

func printElements(c map[string]string) {
	for i, color := range c {
		fmt.Println("Description of the color", color, " key is ", i)
	}
}
