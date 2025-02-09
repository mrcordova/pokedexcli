package main

import "fmt"


func commandPokedex(config *config, args ...string) error  {
	
	fmt.Println("Your Pokedex:")
	for pokemonName := range config.caughtPokemon {
		fmt.Printf("- %v\n", pokemonName)
	}
	return nil
}