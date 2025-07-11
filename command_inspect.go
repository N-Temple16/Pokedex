package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, pokemon []string) error {
	if len(pokemon) == 0 {
        return errors.New("you must provide the name of a pokemon to inspect")
	}

	pokedexEntry, ok := caughtPokemon[pokemon[0]]
	if !ok {
        return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokedexEntry.Name)
	fmt.Printf("Height: %d\n", pokedexEntry.Height)
	fmt.Printf("Weight: %d\n", pokedexEntry.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokedexEntry.Stats {
		fmt.Printf("  -%s: %d\n", stats.StatName.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokedexEntry.Types {
		fmt.Printf("  - %s\n", types.TypeName.Name)
	}

	return nil
}