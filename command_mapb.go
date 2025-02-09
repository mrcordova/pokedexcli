package main

import "fmt"

func commandMapB(config *Config) error {
	previous := config.Previous
	endpoint_url := POKEDEX_URL + "/location-area"
	if previous != "" {
		endpoint_url = previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	
	

	return printMapResults(endpoint_url, config)
}