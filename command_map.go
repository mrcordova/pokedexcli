package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	next := config.Next
	endpoint_url := POKEDEX_URL + "/location-area"
	if next != "" {
		endpoint_url = next
	} 
	
	

	return printMapResults(endpoint_url, config)
}

func printMapResults(url string, config *Config)  error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("status: %v %v", res.StatusCode, res.Status )
	}
	locationResponse := PokedexResponse{}
	err = json.Unmarshal(body, &locationResponse)

	if err != nil {
		return err
	}

	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous
	results := locationResponse.Results

	for _, result := range results {
		fmt.Println(result.Name)

	}
	return nil
}