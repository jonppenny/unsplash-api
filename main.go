package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jonppenny/unsplash-api/endpoints"
	"github.com/jonppenny/unsplash-api/user"
)

func main() {
	config := flag.String("config", "config.toml", "Config file for default settings.")
	query := flag.String("query", "", "Search query")
	perPage := flag.String("per_page", "1", "Results per page, defaults to one")
	flag.Parse()

	credentials, err := user.GetUserCredentials(config)
	if err != nil {
		return
	}

	params := make(map[string]string)
	params["query"] = *query
	params["per_page"] = *perPage

	data, err := endpoints.Get("https://api.unsplash.com/search/photos", credentials, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(data)
}
