package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(config *Config, location []string) error {
	var body []byte
	var err error

	if len(location) == 0 {
        return errors.New("you must provide a location to explore")
	}

	exploreLocation := "https://pokeapi.co/api/v2/location-area/" + location[0]

	data, ok := cache.Get(exploreLocation)
	if ok {
		body = data
	} else {
		res, err := http.Get(exploreLocation)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(exploreLocation, bodyBytes)
		body = bodyBytes
	}

	pokemon := SpecifiedLocation{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location[0])
	fmt.Println("Found Pokemon:")

	for _, poke := range pokemon.Encounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}

    return nil
}