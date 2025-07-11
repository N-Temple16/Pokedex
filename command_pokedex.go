package main

import (
	"fmt"
)

func commandPokedex(config *Config, strings []string) error {
	fmt.Println("Your Pokedex:")
	for _, names := range caughtPokemon {
		fmt.Printf(" - %s\n", names.Name)
	}

	return nil
}