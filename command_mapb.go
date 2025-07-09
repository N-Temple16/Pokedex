package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(config *Config) error {
	var body []byte
	var err error

	if config.Previous == "" {
        fmt.Println("you're on the first page")
        return nil
    }

	data, ok := cache.Get(config.Previous)
	if ok {
		body = data
	} else {
		res, err := http.Get(config.Previous)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(config.Previous, bodyBytes)
		body = bodyBytes
	}

	locations := LocationResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	config.Next = locations.Next
	if locations.Previous != nil {
		config.Previous = *locations.Previous
	} else {
		config.Previous = ""
	}	

    return nil
}