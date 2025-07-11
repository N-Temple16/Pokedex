package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(config *Config, pokemon []string) error {
	var body []byte
	var err error

	if len(pokemon) == 0 {
        return errors.New("you must provide a pokemon to catch")
	}

	catchablePokemon := "https://pokeapi.co/api/v2/pokemon/" + pokemon[0]

	data, ok := cache.Get(catchablePokemon)
	if ok {
		body = data
	} else {
		res, err := http.Get(catchablePokemon)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(catchablePokemon, bodyBytes)
		body = bodyBytes
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon[0])

	catchRate := max(10, 100 - pokemonData.BaseExperience)
	randomNum := rand.Intn(100)

	if randomNum < catchRate {
		fmt.Printf("%s was caught!\n", pokemon[0])
		fmt.Println("You may now inspect it with the inspect command.")
		caughtPokemon[pokemon[0]] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemon[0])
	}
	
	return nil
}