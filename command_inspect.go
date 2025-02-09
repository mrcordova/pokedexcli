package main

import (
	"errors"
	"fmt"
)


func commandInspect(config *config, args ...string) error  {
	if len(args) != 1 {
		return errors.New("please provide pokemon name")
	}
	pokemonName := args[0]
	if pokemon, ok := config.caughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		// fmt.Printf(": %s", pokemon.Stats)
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		for _, t := range pokemon.Types {
			fmt.Printf(" - %v\n", t.Type.Name)
		}

	} else {
		fmt.Println("you have not caught this pokemon")
	}
	return nil
}