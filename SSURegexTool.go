package main

import (
	"fmt"
	"regexp"
)

func main() {
	var entry string
	fmt.Println("********************************************************")
	fmt.Println("RegEx Validation Tool for Servicing Stack Updates (SSUs")
	fmt.Println("********************************************************")
	fmt.Println("This program allows you to enter a string which will be validated against a RegEx pattern written to catch SSUs.")
	fmt.Println("\nAn example of which is: '2022-09 Servicing Stack Update for Windows 8.1 for x86-based Systems (KB5017398)'")
	fmt.Println("\nIf the program returns 'true', it means the string has been caught by the RegEx Pattern")
	fmt.Println("\nEnter string: ")
	fmt.Scan(&entry)

	// Backticks mean you do not need to escape characters
	ssuRegexPattern, match := regexp.MatchString(`([0-9].[0-9].-.\w+.\bServicing Stack Update\b.*)`, entry)

	fmt.Println(ssuRegexPattern, match)
}
