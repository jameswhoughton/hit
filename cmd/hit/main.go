package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
)

func verbIsValid(verb string) bool {
	allowedVerbs := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	}

	return slices.Contains(allowedVerbs, verb)
}

type headers []string

func (h *headers) String() string {
	return strings.Join(*h, ", ")
}

func (h *headers) Set(value string) error {
	*h = append(*h, value)

	return nil
}

func main() {

	var url string

	// Default verb
	verb := "GET"

	var headers headers

	flag.Var(&headers, "header", "")

	flag.Parse()

	args := flag.Args()

	switch len(args) {
	case 0:
		log.Println("at least one argument is required")

		os.Exit(1)
	case 1:
		url = args[0]
	case 2:
		verb = args[0]
		url = args[1]
	}

	if !verbIsValid(verb) {
		log.Printf("verb %s is invalid\n", verb)

		os.Exit(1)
	}

	fmt.Printf("running: %s %s\n", verb, url)

	client := &http.Client{}

	req, err := http.NewRequest(verb, url, nil)

	if err != nil {
		log.Printf("error creating request: %v\n", err)

		os.Exit(1)
	}

	for _, header := range headers {
		parts := strings.Split(header, ":")

		req.Header.Add(parts[0], parts[1])
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)

		os.Exit(1)
	}

	fmt.Print(resp)

	os.Exit(0)
}
