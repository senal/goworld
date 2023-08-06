package main

import (
	"fmt"
	"os"
)

/*
Create a programe that reads the content of a text file then prints its contents to the terminal.

The file to open should be provided as an argument to the program when it is executed at the terminal. For example, 'go run main.go myFile.txt'.
should open up the myFile.txt file

To read in the arguments provided to a program, you can reference the variable 'os.Args', which is a slice of type string

To open a file, check out the documentation for the 'Open' function in the 'os' package - https://pkg.go.dev/os#Open

What interfaces does the 'File' type implement?

If the 'File' type implements the 'Reader' interface, you might be able to reuse that io.Copy function!
*/

func main() {
	// open a file io.Open() which returns a *File, error
	// pass the *File to
	bs, err := os.ReadFile(os.Args[1])

	if err != nil {
		os.Exit(-1)
	}

	fmt.Println(string(bs))
}
