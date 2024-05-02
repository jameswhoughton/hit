package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)

func verbIsValid(verb string) bool {
	allowedVerbs := []string{"GET", "POST", "PUT", "DELETE"}

	return slices.Contains(allowedVerbs, verb)
}

func main() {
	args := os.Args[1:]

	var url string

	// Default verb
	verb := "GET"

	switch len(args) {
	case 0:
		log.Println("at least one argument is required")

		os.Exit(1)
	case 1:
		url = args[0]
	case 2:
		verb = args[0]
		url = args[1]
	default:
		log.Println("too many arguments")

		os.Exit(1)
	}

	if !verbIsValid(verb) {
		log.Printf("verb %s is invalid\n", verb)

		os.Exit(1)
	}

	fmt.Printf("running: %s %s\n", verb, url)
}
